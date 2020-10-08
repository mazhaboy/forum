package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"

	model "../model"
)

var Flag bool
var email string

func login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {

			if r.URL.Path != "/" {
				http.Error(w, "Error 404", http.StatusNotFound)
			}

			if r.URL.Path == "/" {
				http.ServeFile(w, r, "view/main.html")
			}
		}

		if r.Method == "POST" {
			email = r.FormValue("email")

			password := r.FormValue("password")

			if model.IsValid(email, password) == true {

				cookie, err := r.Cookie(email)
				if err == nil {

					fmt.Println(Flag)
					fmt.Fprintln(w, cookie)
					fmt.Println("est uzhe")
				}

				if err != nil {

					expire := time.Now().Add(24 * time.Hour)

					u1, _ := uuid.NewV4()
					cookie = &http.Cookie{
						Name:     email,
						Value:    u1.String(),
						Expires:  expire,
						Path:     "/",
						HttpOnly: true,
					}

					http.SetCookie(w, cookie)
					fmt.Fprintln(w, cookie)
					if err := model.AddSession(email, cookie.Value); err != nil {
						if err := model.UpdateSession(cookie.Value, cookie.Name); err != nil {
							log.Fatal(err)
						}
					}

					fmt.Println("trueeee")
					fmt.Println("net no ya dobavil")
				}

				return

			}

			http.Error(w, "Invalid email or password", http.StatusNotFound)

		}

	}
}
