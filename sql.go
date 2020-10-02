package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int
	Username string
	Email    string
	Password string
}

func main() {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Successfuly connectted to the database!")

	rusult, err := db.Exec("insert into test (username, email, password) values(?,?,?)", "aibekasas1", "nurbekassa2", "qwertasy")
	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}
	fmt.Println(rusult.LastInsertId)
	fmt.Println("Data has successfuly inserted")

}
