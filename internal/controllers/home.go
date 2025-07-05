package controllers

import (
	"html/template"
	"net/http"

	"github.com/reftch/go-react-template/configs"
)

var homeTmpl = template.Must(template.ParseFiles(
	"./web/templates/layout.html",
	"./web/templates/home.html",
))

func (c *Controller) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}

	type HomeData struct {
		Title       string
		Environment string
	}

	data := HomeData{
		Title:       "Resume Editor :: Home",
		Environment: configs.Envs.Environment,
	}

	if err := homeTmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
