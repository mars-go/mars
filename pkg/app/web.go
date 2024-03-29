package app

import (
	"fmt"
	"github.com/foolin/goview"
	"html/template"
	"net/http"
	"time"
)

func Run() {

	gvFront := goview.New(goview.Config{
		Root:      "views/web",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := gvFront.Render(w, http.StatusOK, "index", goview.M{
			"title": "Frontend title!",
		})
		if err != nil {
			fmt.Fprintf(w, "Render index error: %v!", err)
		}
	})

	//=========== Backend ===========//

	gvBackend := goview.New(goview.Config{
		Root:      "views/admin",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	http.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		err := gvBackend.Render(w, http.StatusOK, "index", goview.M{
			"title": "Backend title!",
		})
		if err != nil {
			fmt.Fprintf(w, "Render index error: %v!", err)
		}
	})

	fmt.Printf("Server start on :9090")
	http.ListenAndServe(":9090", nil)

}
