// +build !appengine

package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/ebittleman/go-course/blog"
)

var httpFlag *string
var helpFlag *bool

func init() {
	httpFlag = flag.String("http", ":8080", "Define which local port to bind to")
	helpFlag = flag.Bool("help", false, "Flag to print output and exit")
	flag.CommandLine.SetOutput(os.Stderr)
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
