// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/mattn/go-sqlite3"
// )

// type User struct {
// 	Id       int
// 	Username string
// 	Email    string
// 	Password string
// }

// func main() {
// 	db, err := sql.Open("sqlite3", "forum.db")
// 	if err != nil {
// 		fmt.Println("Error")
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	fmt.Println("Successfuly connectted to the database!")

// 	rusult, err := db.Exec("insert into test (username, email, password) values(?,?,?)", "aibekasas1", "nurbekassa2", "qwertasy")
// 	if err != nil {
// 		fmt.Println("Error")
// 		log.Fatal(err)
// 	}
// 	fmt.Println(rusult.LastInsertId)
// 	fmt.Println("Data has successfuly inserted")

// }
package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	// Creating UUID Version 4
	// panic on error
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: %s\n", u2)

}
