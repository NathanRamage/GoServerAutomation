package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

type webData struct {
	Title string
}

var wd = webData {
	Title: "",
}
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("static/*"))
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	wd.Title = "Index"
	fmt.Println(r.URL.Path)
	tpl.ExecuteTemplate(w, "index.gohtml", wd)
}

func main() {
	http.HandleFunc("/", serveIndex)

	log.Fatal(http.ListenAndServe(":80", nil))
}
