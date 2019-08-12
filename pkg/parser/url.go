package parser

import (
	"path/filepath"
	"strings"
)

func Path2URL(input string) string {
	return strings.ReplaceAll(input, string(filepath.Separator), "/")
}
