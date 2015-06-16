package blog

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/ebittleman/go-course/view"
)

type Page struct {
	Vars map[string]interface{}
}

var cachedTemplates *template.Template

func init() {
	http.HandleFunc("/blog", BlogRootHandler)
	cachedTemplates = template.Must(
		template.New("root").Funcs(
			map[string]interface{}{
				"title": strings.Title,
			},
		).ParseGlob("./blog/templates/*.html"),
	)
}

func BlogRootHandler(resp http.ResponseWriter, req *http.Request) {

	layout := view.NewViewModel(
		"layout.html",
		view.Vars{
			"Title": "rendered output",
		},
	)

	content := view.NewViewModel("article.html", nil)

	layout.AddChild(content, "Content")

	view.RenderViewModel(resp, cachedTemplates, layout)
}
