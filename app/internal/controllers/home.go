package controllers

import (
	"html/template"
	"net/http"
)

var homeTmpl = template.Must(template.ParseFiles(
	"./app/web/templates/layout.html",
	"./app/web/templates/home.html",
))

func (c *Controller) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}

	type HomeData struct {
		Title string
	}

	data := HomeData{
		Title: "Resume Editor :: Home",
	}

	if err := homeTmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
