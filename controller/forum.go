package controller

import (
	"net/http"
)

func forum() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if r.URL.Path != "/" {
				http.Error(w, "Error 404", http.StatusNotFound)
			} else {

				http.ServeFile(w, r, "view/templates.html")

			}
		}
	}
}
