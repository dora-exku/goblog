package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"io"
	"path/filepath"
	"strings"
	"text/template"
)

type D map[string]interface{}

func Render(w io.Writer, data interface{}, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)

}

func RenderSimple(w io.Writer, data interface{}, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

func RenderTemplate(w io.Writer, name string, data interface{}, tplFiles ...string) {
	viewDir := "resources/views/"
	for i := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(tplFiles[i], ".", "/", -1) + ".gohtml"
	}
	files, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)
	newFiles := append(files, tplFiles...)
	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(newFiles...)
	logger.LogError(err)
	tmpl.ExecuteTemplate(w, name, data)
}
