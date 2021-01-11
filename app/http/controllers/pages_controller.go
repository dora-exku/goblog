package controllers

import (
	"fmt"
	"goblog/pkg/view"
	"net/http"
)

type PagesController struct{}

func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "欢迎来到 goblog")
}

func (*PagesController) About(w http.ResponseWriter, r *http.Request) {

	view.Render(w, ArticlesFormData{}, "pages.about")
}

func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 not found")
}
