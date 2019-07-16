package main

import (
	"github.com/mars-go/mars/internal/db"
	model2 "github.com/mars-go/mars/internal/model"
)

func main()  {
	println("Hello world")

	err := db.Init("mars.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	model := &model2.Page{
		Metas:
	}
	
	db.Create("pages", )
	
}
