package blog

import (
	"html"
	"log"
	"net/http"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/ebittleman/go-course/jsondb"
	"github.com/ebittleman/go-course/view"
)

const (
	TableName      = "articles"
	LayoutTemplate = "layout.html"
)

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
				"empty": view.Empty,
				"date":  view.DateFormat("01/02/06, 03:04PM"),
				"eq":    view.Equals,
			},
		).ParseGlob("./blog/templates/*.html"),
	)

	http.HandleFunc("/blog", RootBlogHandler)
	http.HandleFunc("/blog/new", NewBlogHandler)
	http.HandleFunc("/blog/create", CreateBlogHandler)
}

func RootBlogHandler(resp http.ResponseWriter, req *http.Request) {
	sort.Sort(ByDate(articles))
	content := view.NewViewModel(
		"summary.html",
		view.Vars{
			"Articles": articles,
		},
	)

	view.RenderViewModel(
		resp,
		cachedTemplates,
		newLayout("My Articles", "home").
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
		newLayout("New Articles", "new").
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
		Created: time.Now(),
	}

	articles = append(articles, article)

	resp.Write([]byte("\"ok\""))

	writeArticles()
}

func newLayout(title string, page string) *view.ViewModel {
	return view.NewViewModel(
		LayoutTemplate,
		view.Vars{
			"Title": title,
			"Page":  page,
		},
	)
}

func writeArticles() {
	jsondb.WriteTable(TableName, articles)
}
