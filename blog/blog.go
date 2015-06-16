package blog

import (
	"html"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/ebittleman/go-course/jsondb"
	"github.com/ebittleman/go-course/view"
)

const (
	TableName      = "articles"
	LayoutTemplate = "layout.html"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var cachedTemplates *template.Template

var articles []*Article

func init() {

	articles = make([]*Article, 0, 50)

	err := jsondb.LoadTable(TableName, &articles)

	if err != nil {
		log.Println(err)
	}

	cachedTemplates = template.Must(
		template.New("root").Funcs(
			map[string]interface{}{
				"title": strings.Title,
				"html":  html.EscapeString,
				"if":    view.If,
			},
		).ParseGlob("./blog/templates/*.html"),
	)

	http.HandleFunc("/blog", RootBlogHandler)
	http.HandleFunc("/blog/new", NewBlogHandler)
	http.HandleFunc("/blog/create", CreateBlogHandler)
}

func RootBlogHandler(resp http.ResponseWriter, req *http.Request) {
	content := view.NewViewModel(
		"summary.html",
		view.Vars{
			"Articles": articles,
		},
	)

	view.RenderViewModel(
		resp,
		cachedTemplates,
		newLayout("My Articles").
			AddChild(content, "Content"),
	)

	writeArticles()
}

func NewBlogHandler(resp http.ResponseWriter, req *http.Request) {
	content := view.NewViewModel(
		"creator.html",
		view.Vars{
			"Articles": articles,
		},
	)

	footerScript := view.NewViewModel(
		"creatorScript",
		nil,
	)

	view.RenderViewModel(
		resp,
		cachedTemplates,
		newLayout("My Articles").
			AddChild(content, "Content").
			AddChild(footerScript, "Scripts"),
	)

	writeArticles()
}

func CreateBlogHandler(resp http.ResponseWriter, req *http.Request) {

	title := req.FormValue("title")
	content := req.FormValue("content")

	article := &Article{
		Title:   title,
		Content: content,
	}

	articles = append(articles, article)

	resp.Write([]byte("\"ok\""))

	writeArticles()
}

func newLayout(title string) *view.ViewModel {
	return view.NewViewModel(
		LayoutTemplate,
		view.Vars{
			"Title": "rendered output",
		},
	)
}

func writeArticles() {
	jsondb.WriteTable(TableName, articles)
}
