package model

import (
	"github.com/mars-go/mars/internal/db"
)

type MetaKey string

const (
	MetaTitle       MetaKey = "title"
	MetaSlug        MetaKey = "slug"
	MetaDate        MetaKey = "date"
	MetaKeywords    MetaKey = "keywords"
	MetaDescription MetaKey = "description"
	MetaDraft       MetaKey = "draft"
	MetaSummary     MetaKey = "summary"
	MetaLayout      MetaKey = "layout"
	MetaType        MetaKey = "type"
	MetaUrl         MetaKey = "url"
	MetaWeight      MetaKey = "weight"
	//publishDate
	//expiryDate
	//taxonomies
)

type Page struct {
	db.Model
	Metas   map[string]interface{}
	Content string
}
