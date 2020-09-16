package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Connected to the database")

	db.Exec("INSERT INTO users (username,email,password) VALUES (?, ?, ? )", "wqqwe", "qweqwe@MAIL.RU", "accw")

	fmt.Println("Data is inserted")

}
