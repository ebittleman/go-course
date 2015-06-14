package cart

import (
	"io"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

const (
	baseDir = "./cart"
)

var productTemplate *template.Template

func init() {
	http.HandleFunc("/cart", HandleRequest)
}

func HandleRequest(resp http.ResponseWriter, req *http.Request) {

	productTemplate = template.Must(template.ParseFiles(filepath.Join(baseDir, "templates", "product.html")))
	err := productTemplate.Execute(resp, nil)

	if err != nil {
		log.Println(err)
	}
}

func HandleCreateProduct(resp http.ResponseWriter, req *http.Request) {
	io.Copy()
}
