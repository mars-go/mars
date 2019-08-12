package parser

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"
)

func Markdown2HTML(input []byte) []byte {
	unsafe := blackfriday.Run(input)
	return bluemonday.UGCPolicy().SanitizeBytes(unsafe)
}

func Markdown2HTMLString(input []byte) template.HTML {
	return template.HTML(Markdown2HTML(input))
}
