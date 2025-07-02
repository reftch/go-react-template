package controllers

import "net/http"

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) GET(path string, f http.HandlerFunc) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		f(w, r)
	})
}
