package main

import (
	"log"
	"net/http"

	controller "./controller"
)

func main() {
	mux := controller.Register()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
