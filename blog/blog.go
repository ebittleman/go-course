package blog

import (
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/ebittleman/go-course/jsondb"
	"github.com/ebittleman/go-course/view"
)

const (
	TableName      = "articles"
	LayoutTemplate = "layout.html"
)

var articles []*Article

func init() {

	articles = make([]*Article, 0, 50)

	err := jsondb.LoadTable(TableName, &articles)

	if err != nil {
		log.Println(err)
	}

	view.RegisterGlob("./blog/templates/*.html")

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
		newLayout("My Articles", "home").
			AddChild(content, "Content"),
	)
}

func NewBlogHandler(resp http.ResponseWriter, req *http.Request) {
	content := view.NewViewModel(
		"creator.html",
		nil,
	)

	footerScript := view.NewViewModel(
		"creatorScript",
		nil,
	)

	view.RenderViewModel(
		resp,
		newLayout("New Articles", "new").
			AddChild(content, "Content").
			AddChild(footerScript, "Scripts"),
	)
}

func CreateBlogHandler(resp http.ResponseWriter, req *http.Request) {
	defer writeArticles()

	title := req.FormValue("title")
	content := req.FormValue("content")

	article := &Article{
		Title:   title,
		Content: content,
		Created: time.Now(),
	}

	articles = append(articles, article)

	resp.Write([]byte("\"ok\""))

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
