package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	model "../model"
)

var Flag bool

func postsandlikes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates := template.Must(template.ParseGlob("view/*.html"))

		posts := model.GetPosts()

		if r.Method == "GET" {
			var User_ID int
			var UserName string

			cookie, err := r.Cookie("session")

			if err != nil {
				Flag = false
			} else {
				Flag = model.IsUserValid(cookie.Value)
				User_ID, UserName = model.GetUserIDbySession(cookie.Value)
				fmt.Println(UserName)
				fmt.Println(User_ID)
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
			cookie, _ := r.Cookie("session")
			User_ID, UserName := model.GetUserIDbySession(cookie.Value)
			fmt.Println(UserName)
			fmt.Println(User_ID)

			post := r.FormValue("post")
			fmt.Println(post)

			if len(post) > 0 {
				if err := model.AddPost(User_ID, post, UserName); err != nil {
					log.Fatal(err)
				}

			}
			post = ""

			Post_IDL, _ := strconv.Atoi(r.FormValue("Post_ID"))

			comment := r.FormValue("comment")
			Post_IDC, _ := strconv.Atoi(r.FormValue("Post_IDC"))

			fmt.Println(Post_IDC)

			if Post_IDL != 0 {

				if model.IsLiked(User_ID, Post_IDL) == false {
					if err := model.AddLike(User_ID, Post_IDL); err != nil {
						log.Fatal(err)
					}
				} else {

					fmt.Println("Like is deleted")

				}
			}
			if Post_IDC != 0 {
				if err := model.AddComment(User_ID, Post_IDC, comment, UserName); err != nil {
					log.Fatal(err)
				}
				fmt.Println(comment)
				fmt.Println(Post_IDC)
			}
			fmt.Println("tut")
			fmt.Println(Post_IDL)

			http.Redirect(w, r, "/posts", 302)

		}
	}
}
