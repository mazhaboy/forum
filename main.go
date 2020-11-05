package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mazhaboy/forum/tree/master/controller"
	"github.com/mazhaboy/forum/tree/master/model"
	// controller "./controller"
	// model "./model"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":8282", mux))
}
