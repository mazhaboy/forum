package controller

import "net/http"

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/registration", registration())
	mux.HandleFunc("/", login())
	return mux
}
