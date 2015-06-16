package view

import (
	"bytes"
	"io"
	"text/template"
)

func RenderViewModel(w io.Writer, templates *template.Template, model *ViewModel) {
	for captureTo, children := range model.Children {
		buffer := &bytes.Buffer{}
		for _, child := range children {
			RenderViewModel(buffer, templates, child)
		}

		model.Vars[captureTo] = buffer.String()
	}

	templates.ExecuteTemplate(w, model.TemplateName, model.Vars)
}
