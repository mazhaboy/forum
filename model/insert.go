package model

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Insert(a, b, c string) error { //Y

	hashed, _ := bcrypt.GenerateFromPassword([]byte(c), bcrypt.DefaultCost)
	password := string(hashed)
	_, err := con.Exec("INSERT INTO User (Email,UserName,Password) VALUES(?,?,?)", b, a, password)
	if err != nil {
		return err
	}
	fmt.Println("Data is inserted")
	return nil

}
func AddSession(a, b string) error { //Y

	_, err := con.Exec("INSERT INTO Session (Email,Session_ID) VALUES(?,?)", a, b)
	if err != nil {
		fmt.Println("Errrrrrrrrr")
		return err
	}
	fmt.Println("Data is inserted")
	return nil

}

func AddPost(a int, b string, UserName string) error {

	t := time.Now()
	fmt.Println(t.String())
	T := t.Format("2006-01-02 15:04:05")
	date := T[0:11]
	time := T[11:]
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	_, err := con.Exec("INSERT INTO Post (User_ID, Post_body, Post_date,Post_time, UserName) VALUES(?,?,?,?,?)", a, b, date, time, UserName)
	if err != nil {
		fmt.Println("Errrrrrrrrr")
		return err
	}
	fmt.Println("Data is inserted")
	return nil

}

func AddLike(a int, b int) error {

	_, err := con.Exec("INSERT INTO Like (User_ID, Post_ID) VALUES(?,?)", a, b)
	if err != nil {
		fmt.Println("Errrrrrrrrr")
		return err
	}
	fmt.Println("Data is inserted")
	return nil

}
func AddComment(a int, b int, comment string, UserName string) error {

	_, err := con.Exec("INSERT INTO Comment (User_ID, Post_ID, Comment, UserName) VALUES(?,?,?,?)", a, b, comment, UserName)
	if err != nil {
		fmt.Println("Errrrrrrrrr")
		return err
	}
	fmt.Println("Data is inserted")
	return nil

}
