// +build appengine

package main

import (
	"github.com/ebittleman/go-course/view"
	_ "golang.org/x/tools/cmd/present"
)

func init() {
	// register user interface elements
	view.RegisterGlob("./templates/*.html")
}
