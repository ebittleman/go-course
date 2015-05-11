package guestbook

import "time"

type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}
