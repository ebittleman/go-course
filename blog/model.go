package blog

import "time"

type Post struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

type ByDate []*Post

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Created.After(a[j].Created) }

type PostTable struct {
	NextId int     `json:"nextId"`
	Posts  []*Post `json:"posts"`
}

func NewPostTable() *PostTable {
	return &PostTable{NextId: 1, Posts: make([]*Post, 0, 10)}
}

func FindPostById(posts []*Post, id int) (*Post, int) {
	for index, post := range posts {
		if post.Id == id {
			return post, index
		}
	}

	return nil, -1
}
