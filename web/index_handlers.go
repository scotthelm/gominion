package web

import (
	"github.com/GeertJohan/go.rice"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpls, err := rice.FindBox("../templates")
	if err != nil {
		log.Fatal(err)
	}
	t, err := tmpls.String("index.html")
	if err != nil {
		log.Panic(err)
	}
	tmpl, err := template.New("Index").Parse(t)
	if err != nil {
		log.Panic(err)
	}
	tmpl.Execute(w, nil)
}
