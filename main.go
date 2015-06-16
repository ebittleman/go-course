// +build !appengine

package main

import (
	"log"
	"net/http"

	_ "github.com/ebittleman/go-course/blog"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
