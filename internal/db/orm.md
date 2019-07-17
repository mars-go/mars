```go

package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	TagDB       = "db"
	TagSplitSep = ";"
	TagKVSep    = ":"
	TagPK       = "pk"
	TagAI       = "ai"
	IDField     = "ID"
)

type FieldInfo struct {
	Field reflect.StructField
	Tags  map[string]string
}

func Save(bucket string, v interface{}) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	// Retrieve the users bucket.
	// This should be created when the DB is first opened.
	bk, err := tx.CreateBucketIfNotExists([]byte(bucket))
	if err != nil {
		return err
	}
	pkFieldInfo, foundPkField := findPkField(v)
	if !foundPkField {
		return errors.New("not found pk field")
	}

	vf := reflect.ValueOf(v).Elem()
	value := vf.FieldByName(pkFieldInfo.Field.Name)
	key := value.String()
	if _, ok := pkFieldInfo.Tags[TagAI]; ok {
		if value.IsValid() {
			key = fmt.Sprintf("%v", value.Interface())
		} else {
			if !value.CanSet() {
				return fmt.Errorf("can't set pk field: %v", pkFieldInfo.Field.Name)
			}
			id, err := bk.NextSequence()
			if err != nil {
				return err
			}
			key = strconv.FormatUint(id, 10)
			switch value.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				value.SetInt(int64(id))
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				value.SetUint(id)
			case reflect.String:
				value.SetString(fmt.Sprintf("%v", id))
			default:
				return fmt.Errorf("field: %v not can't set kind: %v", value.Kind())
			}
		}
	}

	// Marshal user data into bytes.
	buf, err := json.Marshal(v)
	if err != nil {
		return err
	}

	// Persist bytes to users bucket.
	return bk.Put([]byte(key), buf)
}


func findPkField(v interface{}) (*FieldInfo, bool) {
	tf := reflect.TypeOf(v)
	var fieldInfo FieldInfo
	foundPkField := false
	for i := 0; i < tf.NumField(); i++ {
		field := tf.Field(i)
		kvTags := parseDBTags(field)
		if _, ok := kvTags[TagPK]; ok {
			fieldInfo.Field = field
			fieldInfo.Tags = kvTags
			foundPkField = true
			break
		}
	}
	if !foundPkField {
		idField, ok := tf.FieldByName(IDField)
		if !ok {
			return nil, false
		}
		fieldInfo.Field = idField
		fieldInfo.Tags = parseDBTags(idField)
	}
	return &fieldInfo, true
}

func parseDBTags(field reflect.StructField) map[string]string {
	//pk=primary key
	//ai=auto increment
	//db:"PK;AI"
	tags := strings.Split(field.Tag.Get(TagDB), TagSplitSep)
	kvTags := make(map[string]string, len(tags))
	for _, tag := range tags {
		tagArr := strings.Split(tag, TagKVSep)
		tagLen := len(tagArr)
		switch tagLen {
		case 1:
			kvTags[strings.ToLower(strings.TrimSpace(tagArr[0]))] = ""
		case 2:
			kvTags[strings.ToLower(strings.TrimSpace(tagArr[0]))] = strings.TrimSpace(tagArr[2])
		}
	}
	return kvTags
}


```
