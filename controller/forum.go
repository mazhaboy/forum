package controller

import (
	"net/http"

	model "../model"
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
		if r.Method == "POST" {

			username := r.FormValue("username")
			email := r.FormValue("email")
			password := r.FormValue("password")

			if err := model.Insert(username, email, password); err != nil {
				http.Error(w, "Error 400", http.StatusBadRequest)
			}
			http.ServeFile(w, r, "view/templates.html")

		}
	}
}
