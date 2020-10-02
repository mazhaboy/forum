package controller

import (
	"fmt"
	"net/http"

	model "../model"
)

func login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if r.URL.Path != "/" {
				http.Error(w, "Error 404", http.StatusNotFound)
			} else {
				http.ServeFile(w, r, "view/main.html")
			}
		}
		if r.Method == "POST" {
			email := r.FormValue("email")
			password := r.FormValue("password")
			if model.IsValid(email, password) == true {
				fmt.Println("User is valid")
				http.ServeFile(w, r, "view/main.html")
				return
			}
			http.Error(w, "Invalid email or password", http.StatusNotFound)

		}
	}
}
