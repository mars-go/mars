package model

import (
	"github.com/mars-go/mars/pkg/db"
	"time"
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

type Post struct {
	db.Model
	CategoryID uint64                 //分类
	Name       string                 //文件名
	Title      string                 //标题
	Date       time.Time              //日期
	Summary    string                 //描述
	Slug       string                 //路由标识
	URL        string                 //URL路由
	weight     int                    //权重，排序用
	Tags       []string               //标签
	Matters    map[string]interface{} //属性
	Content    string                 //内容
	Draft      bool                   //是否是草稿
}
