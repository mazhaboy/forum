package controller

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"

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

				cookie, err := r.Cookie(email)
				if err != nil {
					u1, _ := uuid.NewV4()
					cookie = &http.Cookie{
						Name:  email,
						Value: u1.String(),
					}
					http.SetCookie(w, cookie)
				}
				fmt.Println(cookie)
				return

			}
			http.Error(w, "Invalid email or password", http.StatusNotFound)

		}
	}
}
