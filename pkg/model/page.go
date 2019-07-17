package model

import (
	"github.com/mars-go/mars/pkg/db"
)

type MatterKey string

const (
	MatterTitle       MatterKey = "title"
	MatterSlug        MatterKey = "slug"
	MatterDate        MatterKey = "date"
	MatterKeywords    MatterKey = "keywords"
	MatterDescription MatterKey = "description"
	MatterDraft       MatterKey = "draft"
	MatterSummary     MatterKey = "summary"
	MatterLayout      MatterKey = "layout"
	MatterType        MatterKey = "type"
	MatterUrl         MatterKey = "url"
	MatterWeight      MatterKey = "weight"
	//publishDate
	//expiryDate
	//taxonomies

	DefaultIndexName = "_index"
)

type Page struct {
	db.Model
	Path    string                 //路径
	Name    string                 //文件名
	Matters map[string]interface{} //属性
	Content string                 //内容
}
