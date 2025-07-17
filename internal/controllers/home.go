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

	// anonymous struct to hold data for the template
	data := struct {
		Title       string
		Environment string
	}{
		Title:       "Resume Editor :: Home",
		Environment: configs.Envs.Environment,
	}

	if err := homeTmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
