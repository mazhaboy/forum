package controller

import (
	"fmt"
	"net/http"
)

func postsandlikes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			fmt.Fprintln(w, "hello posts")

		}

		if r.Method == "POST" {

		}
	}
}
