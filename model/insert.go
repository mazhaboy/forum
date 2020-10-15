package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Insert(a, b, c string) error {

	hashed, _ := bcrypt.GenerateFromPassword([]byte(c), bcrypt.DefaultCost)
	password := string(hashed)
	_, err := con.Exec("INSERT INTO test (username,email,password) VALUES(?,?,?)", a, b, password)
	if err != nil {
		return err
	}
	fmt.Println("Data is inserted")
	return nil

}
func AddSession(a, b string) error {

	_, err := con.Exec("INSERT INTO  post (Email,SessionID) VALUES(?,?)", a, b)
	if err != nil {
		fmt.Println("Errrrrrrrrr")
		return err
	}
	fmt.Println("Data is inserted")
	return nil

}
func AddPost(a, b string) error {

	_, err := con.Exec("INSERT INTO  pl (PostID,Email, Post, Like) VALUES(?,?,?,?)", 10, a, b, 0)
	if err != nil {
		fmt.Println("Errrrrrrrrr")
		return err
	}
	fmt.Println("Data is inserted")
	return nil

}
