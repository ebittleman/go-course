// +build !appengine

package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/ebittleman/go-course/blog"
	_ "github.com/ebittleman/go-course/calc"
	"github.com/ebittleman/go-course/view"
)

var httpFlag *string
var helpFlag *bool

func init() {
	httpFlag = flag.String("http", ":8080", "Define which local port to bind to")
	helpFlag = flag.Bool("help", false, "Flag to print output and exit")
	flag.CommandLine.SetOutput(os.Stderr)

	// register user interface elements
	view.RegisterGlob("./templates/*.html")

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		view.RenderViewModel(
			resp,
			view.NewViewModel(
				"layout.html",
				view.Vars{
					"Title":   "Code Camp 2015",
					"Page":    "home",
					"Content": "<h1>Welcome!</h1>",
				},
			),
		)
	})
}

func main() {
	flag.Parse()

	if *helpFlag {
		flag.PrintDefaults()
		os.Exit(0)
	}

	log.Printf("Running Web Server on %s\n", *httpFlag)
	log.Fatal(http.ListenAndServe(*httpFlag, nil))
}
