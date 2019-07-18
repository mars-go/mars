package model

import (
	"github.com/mars-go/mars/pkg/db"
)

type Category struct {
	db.Model
	Slug string
	Name string
}
