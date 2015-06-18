package calc

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ebittleman/go-course/view"
)

func init() {
	view.RegisterGlob("./calc/templates/*.html")

	http.HandleFunc("/calc", CalcPageHandler)
	http.HandleFunc("/calc/add", HandleAdd)
	http.HandleFunc("/calc/subtract", HandleSubtract)
	http.HandleFunc("/calc/multiply", HandleMultiply)
	http.HandleFunc("/calc/divide", HandleDivide)
}

const (
	LayoutTemplate = "layout.html"
)

type Result struct {
	Answer int `json:"answer"`
}

type Operation func(int, int) int

func CalcPageHandler(resp http.ResponseWriter, req *http.Request) {
	content := view.NewViewModel(
		"calculator.html",
		nil,
	)

	footerScript := view.NewViewModel(
		"calculatorScript",
		nil,
	)

	view.RenderViewModel(
		resp,
		newLayout("Calculator", "calc").
			AddChild(content, "Content").
			AddChild(footerScript, "Scripts"),
	)
}

func HandleAdd(resp http.ResponseWriter, req *http.Request) {
	handleRequest(resp, req, Add)

}

func HandleSubtract(resp http.ResponseWriter, req *http.Request) {
	handleRequest(resp, req, Subtract)
}

func HandleMultiply(resp http.ResponseWriter, req *http.Request) {
	handleRequest(resp, req, Multiply)
}

func HandleDivide(resp http.ResponseWriter, req *http.Request) {
	handleRequest(resp, req, Divide)
}

func handleRequest(resp http.ResponseWriter, req *http.Request, operation Operation) {
	a, b, err := parseRequest(req)

	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	var result *Result = &Result{
		Answer: operation(a, b),
	}

	data, err := json.Marshal(result)

	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(data)
}

func parseRequest(req *http.Request) (int, int, error) {
	aStr := req.FormValue("a")
	bStr := req.FormValue("b")

	a, err := strconv.ParseInt(aStr, 10, 64)

	if err != nil {
		return 0, 0, err
	}

	b, err := strconv.ParseInt(bStr, 10, 64)

	if err != nil {
		return 0, 0, err
	}

	return int(a), int(b), nil
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
