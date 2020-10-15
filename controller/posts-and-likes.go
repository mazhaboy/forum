package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	model "../model"
)

var Flag bool
var Username string

func postsandlikes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates := template.Must(template.ParseGlob("view/*.html"))

		posts := model.GetPosts()

		if r.Method == "GET" {

			cookie, err := r.Cookie("session")

			if err != nil {
				Flag = false
			} else {
				Flag = model.IsUserValid(cookie.Value)
				Username = model.GetUsername(cookie.Value)
				fmt.Println(cookie.Value)
			}

			if Flag {
				if err := templates.ExecuteTemplate(w, "posts.html", posts); err != nil {
					http.Error(w, "Internal Server Error!!!\nERROR-500", http.StatusInternalServerError)
					return
				}

			} else {
				if err := templates.ExecuteTemplate(w, "postsonly.html", posts); err != nil {
					http.Error(w, "Internal Server Error!!!\nERROR-500", http.StatusInternalServerError)
					return
				}

			}

		}

		if r.Method == "POST" {

			post := r.FormValue("post")
			fmt.Println(post)
			if len(post) > 0 {
				if err := model.AddPost(Username, post); err != nil {
					log.Fatal(err)
				}
				post = ""
			}
			fmt.Println(post)
			if err := templates.ExecuteTemplate(w, "posts.html", posts); err != nil {
				http.Error(w, "Internal Server Error!!!\nERROR-500", http.StatusInternalServerError)
				return
			}

		}
	}
}
