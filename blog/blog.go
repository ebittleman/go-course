package blog

import (
	"log"
	"net/http"
	"path"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/ebittleman/go-course/jsondb"
	"github.com/ebittleman/go-course/view"
)

const (
	TableName      = "posts"
	LayoutTemplate = "layout.html"
)

var postsTable *PostTable
var lock *sync.Mutex

func init() {

	// initialize data
	postsTable = NewPostTable()
	lock = &sync.Mutex{}

	err := jsondb.LoadTable(TableName, postsTable)

	if err != nil {
		log.Println(err)
	}

	// register user interface elements
	view.RegisterGlob("./blog/templates/*.html")

	// register url paths and handlers
	http.HandleFunc("/blog", ListPostsHandler)
	http.HandleFunc("/blog/new", NewPostFormHandler)
	http.HandleFunc("/blog/create", CreatePostHandler)
	http.HandleFunc("/blog/edit/", EditPostFormHandler)
	http.HandleFunc("/blog/update/", UpdatePostFormHandler)
	http.HandleFunc("/blog/delete/", DeletePostFormHandler)
}

func ListPostsHandler(resp http.ResponseWriter, req *http.Request) {
	sort.Sort(ByDate(postsTable.Posts))

	content := view.NewViewModel(
		"summary.html",
		view.Vars{
			"Posts": postsTable.Posts,
		},
	)

	view.RenderViewModel(
		resp,
		newLayout("My Posts", "posts").
			AddChild(content, "Content"),
	)
}

func NewPostFormHandler(resp http.ResponseWriter, req *http.Request) {
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
		newLayout("New Post", "new").
			AddChild(content, "Content").
			AddChild(footerScript, "Scripts"),
	)
}

func CreatePostHandler(resp http.ResponseWriter, req *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	defer writePosts()

	title := req.FormValue("title")
	content := req.FormValue("content")

	post := &Post{
		Id:      postsTable.NextId,
		Title:   title,
		Content: content,
		Created: time.Now(),
	}

	postsTable.Posts = append(postsTable.Posts, post)
	postsTable.NextId += 1

	resp.Write([]byte("\"ok\""))
}

func EditPostFormHandler(resp http.ResponseWriter, req *http.Request) {

	_, id := path.Split(req.RequestURI)

	postId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	post, _ := FindPostById(postsTable.Posts, int(postId))

	if post == nil {
		http.Error(resp, "Post Not Found", http.StatusNotFound)
		return
	}

	content := view.NewViewModel(
		"creator.html",
		view.Vars{
			"Post": post,
		},
	)

	footerScript := view.NewViewModel(
		"creatorScript",
		nil,
	)

	view.RenderViewModel(
		resp,
		newLayout("New Post", "new").
			AddChild(content, "Content").
			AddChild(footerScript, "Scripts"),
	)
}

func UpdatePostFormHandler(resp http.ResponseWriter, req *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	defer writePosts()

	_, id := path.Split(req.RequestURI)

	postId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	post, _ := FindPostById(postsTable.Posts, int(postId))

	if post == nil {
		http.Error(resp, "Post Not Found", http.StatusNotFound)
		return
	}

	title := req.FormValue("title")
	content := req.FormValue("content")

	post.Title = title
	post.Content = content

	resp.Write([]byte("\"ok\""))
}

func DeletePostFormHandler(resp http.ResponseWriter, req *http.Request) {
	lock.Lock()
	defer http.Redirect(resp, req, "/blog", http.StatusTemporaryRedirect)
	defer lock.Unlock()
	defer writePosts()

	_, id := path.Split(req.RequestURI)

	postId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Println(err)
		return
	}

	post, i := FindPostById(postsTable.Posts, int(postId))

	if post == nil {
		log.Println("Post Not Found")
		return
	}

	postsTable.Posts = append(postsTable.Posts[:i], postsTable.Posts[i+1:]...)
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

func writePosts() {
	jsondb.WriteTable(TableName, postsTable)
}
