package view

import (
	"bytes"
	"io"
	"strings"
	"text/template"
	"time"
)

var defaultFuncs template.FuncMap

var templates *template.Template

func init() {

	defaultFuncs = template.FuncMap{
		"title": strings.Title,
		"date":  DateFormat("01/02/06, 03:04PM", time.Local),
		"empty": Empty,
	}

	templates = template.New("root").Funcs(defaultFuncs)
}

func RegisterHelpers(funcMap template.FuncMap) {
	templates = templates.Funcs(funcMap)
}

func RegisterGlob(pattern string) {
	templates = template.Must(templates.ParseGlob(pattern))
}

func RenderViewModel(w io.Writer, model *ViewModel) {
	for captureTo, children := range model.Children {
		buffer := &bytes.Buffer{}
		for _, child := range children {
			RenderViewModel(buffer, child)
		}

		model.Vars[captureTo] = buffer.String()
	}

	templates.ExecuteTemplate(w, model.TemplateName, model.Vars)
}
