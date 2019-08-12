package parser

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/mars-go/mars/pkg/parser/pageparser"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

func ParseFile(fileName string) (*PageInfo, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("read file %v error: %v", fileName, err)
	}
	defer file.Close()
	return Parse(file)
}

func Parse(reader io.Reader) (*PageInfo, error) {
	var matterType pageparser.ItemType
	var matterValue []byte
	var content []byte

	psr, err := pageparser.Parse(reader, pageparser.Config{})
	if err != nil {
		return nil, err
	}
	iter := psr.Iterator()
	walkFn := func(item pageparser.Item) bool {
		if matterValue != nil {
			// The rest is content.
			content = psr.Input()[item.Pos:]
			// Done
			return false
		} else if item.IsFrontMatter() {
			matterType = item.Type
			matterValue = item.Val
		}
		return true
	}
	iter.PeekWalk(walkFn)

	headers := make(map[string]interface{}, 0)
	switch matterType {
	case pageparser.TypeFrontMatterJSON:
		err = json.Unmarshal(matterValue, &headers)
	case pageparser.TypeFrontMatterTOML:
		err = toml.Unmarshal(matterValue, &headers)
	case pageparser.TypeFrontMatterYAML:
		err = yaml.Unmarshal(matterValue, &headers)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal YAML: %v", err)
		}
		err = yamlKeys(&headers)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("not support matter type: %v", matterType.String())
	}

	info := &PageInfo{
		Matters: headers,
		Content: content,
	}
	return info, nil
}

// To support boolean keys, the YAML package unmarshals maps to
// map[interface{}]interface{}. Here we recurse through the result
// and change all maps to map[string]interface{} like we would've
// gotten from `json`.
func yamlKeys(v interface{}) error {
	var ptr interface{}
	switch v.(type) {
	case *map[string]interface{}:
		ptr = *v.(*map[string]interface{})
	case *interface{}:
		ptr = *v.(*interface{})
	default:
		return fmt.Errorf("unknown type %T in YAML unmarshal", v)
	}

	if mm, changed := stringifyMapKeys(ptr); changed {
		switch v.(type) {
		case *map[string]interface{}:
			*v.(*map[string]interface{}) = mm.(map[string]interface{})
		case *interface{}:
			*v.(*interface{}) = mm
		}
	}
	return nil
}

// stringifyMapKeys recurses into in and changes all instances of
// map[interface{}]interface{} to map[string]interface{}. This is useful to
// work around the impedance mismatch between JSON and YAML unmarshaling that's
// described here: https://github.com/go-yaml/yaml/issues/139
//
// Inspired by https://github.com/stripe/stripe-mock, MIT licensed
func stringifyMapKeys(in interface{}) (interface{}, bool) {

	switch in := in.(type) {
	case []interface{}:
		for i, v := range in {
			if vv, replaced := stringifyMapKeys(v); replaced {
				in[i] = vv
			}
		}
	case map[string]interface{}:
		for k, v := range in {
			if vv, changed := stringifyMapKeys(v); changed {
				in[k] = vv
			}
		}
	case map[interface{}]interface{}:
		res := make(map[string]interface{})
		var (
			ok  bool
			err error
		)
		for k, v := range in {
			var ks string

			if ks, ok = k.(string); !ok {
				ks, err = cast.ToStringE(k)
				if err != nil {
					ks = fmt.Sprintf("%v", k)
				}
			}
			if vv, replaced := stringifyMapKeys(v); replaced {
				res[ks] = vv
			} else {
				res[ks] = v
			}
		}
		return res, true
	}

	return nil, false
}
