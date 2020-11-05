package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	// model "../model"
	// view "../view"
	"github.com/mazhaboy/forum/tree/master/model"
	"github.com/mazhaboy/forum/tree/master/view"
)

var Flag bool
var filter string
var myposts string
var posts []view.Post
var User_ID int
var UserName string

func postsandlikes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates := template.Must(template.ParseGlob("view/*.html"))

		posts = model.GetPosts(filter, User_ID)

		if r.Method == "GET" {

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
					fmt.Println(filter)
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
			User_ID, UserName = model.GetUserIDbySession(cookie.Value)
			fmt.Println(UserName)
			fmt.Println(User_ID)
			myposts := r.FormValue("myposts")
			if myposts == "myposts" {
				filter = myposts
				http.Redirect(w, r, "/posts", 302)
				return
			} else if myposts == "myfavourite" {
				filter = myposts
				http.Redirect(w, r, "/posts", 302)
				return
			}
			fmt.Println(myposts)
			filter = r.FormValue("filter")
			fmt.Println(filter)
			post := r.FormValue("post")
			category := r.FormValue("category")

			fmt.Println(post)
			fmt.Println(category)

			if len(post) > 0 {
				if err := model.AddPost(User_ID, post, UserName, category); err != nil {
					log.Fatal(err)
				}
			}
			post = ""

			Post_IDL, _ := strconv.Atoi(r.FormValue("Post_ID"))

			comment := r.FormValue("comment")
			Post_IDC, _ := strconv.Atoi(r.FormValue("Post_IDC"))
			Comment_ID, _ := strconv.Atoi(r.FormValue("Comment_ID"))

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
			if Comment_ID != 0 {

				if model.IsCommentLiked(Comment_ID, User_ID) == false {
					if err := model.AddCommentLike(Comment_ID, User_ID); err != nil {
						log.Fatal(err)
					}
				} else {

					fmt.Println("Comment Like is deleted")

				}
			}
			Comment_ID = 0
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
