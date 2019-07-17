package app

import (
	"fmt"
	"github.com/foolin/goview"
	"github.com/go-chi/chi"
	"net/http"
)

func Run() {

	r := chi.NewRouter()

	//render index use `index` without `.html` extension, that will render with master layout.
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := goview.Render(w, http.StatusOK, "index", goview.M{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
		if err != nil {
			fmt.Fprintf(w, "Render index error: %v!", err)
		}
	})

	//render page use `page.tpl` with '.html' will only file template without master layout.
	r.Get("/page", func(w http.ResponseWriter, r *http.Request) {
		err := goview.Render(w, http.StatusOK, "page.html", goview.M{"title": "Page file title!!"})
		if err != nil {
			fmt.Fprintf(w, "Render page.html error: %v!", err)
		}
	})

	fmt.Println("Listening and serving HTTP on :9090")
	http.ListenAndServe(":9090", r)

}
