package view

import (
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"io"
	"path/filepath"
	"strings"
	"text/template"
)

type D map[string]interface{}

func Render(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)

}

func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string) {

	data["IsLogined"] = auth.Check()
	data["loginUser"] = auth.User
	data["flash"] = flash.All()

	allFiles := GetTemplateFiles(tplFiles...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(allFiles...)
	logger.LogError(err)
	tmpl.ExecuteTemplate(w, name, data)
}

func GetTemplateFiles(tplFiles ...string) []string {
	viewDir := "resources/views/"
	for i := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(tplFiles[i], ".", "/", -1) + ".gohtml"
	}
	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")

	logger.LogError(err)
	return append(layoutFiles, tplFiles...)
}
