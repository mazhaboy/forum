package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	model "../model"
)

func postsandlikes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates := template.Must(template.ParseGlob("view/*.html"))
		posts := model.GetPosts()
		if r.Method == "GET" {
			if Flag {
				templates.ExecuteTemplate(w, "posts.html", posts)
			} else {
				fmt.Fprintln(w, "Bez est zhe)")
			}

		}

		if r.Method == "POST" {
			post := r.FormValue("post")
			if err := model.AddPost(email, post); err != nil {
				log.Fatal(err)
			}

			if err := templates.ExecuteTemplate(w, "posts.html", posts); err != nil {
				http.Error(w, "Internal Server Error!!!\nERROR-500", http.StatusInternalServerError)
				return
			}

		}
	}
}
