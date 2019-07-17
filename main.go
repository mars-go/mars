package main

import (
	"github.com/mars-go/mars/pkg/app"
	"github.com/mars-go/mars/pkg/db"
	"github.com/mars-go/mars/pkg/model"
	"github.com/mars-go/mars/pkg/utilx"
	"github.com/mars-go/mars/pkg/utilx/logx"
)

func main() {

	//日志
	//logx.Init(false)
	//defer logx.Flush()

	err := db.Init("mars.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	optDB := db.Get()

	page := model.Page{
		Name: "Hello world",
	}

	err = optDB.BucketCreate("pages", &page)
	if err != nil {
		logx.Error(err)
	}
	logx.Infof("page: %#v\n", utilx.JSON(page))

	app.Run()

}
