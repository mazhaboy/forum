package model

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	view "../view"
)

func IsValid(a, b string) bool {

	rows, err := con.Query("select * from test")
	if err != nil {
		log.Fatal(err)
	}
	Users := []view.User{}
	for rows.Next() {
		u := view.User{}
		err := rows.Scan(&u.Id, &u.Email, &u.Username, &u.Password)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		Users = append(Users, u)
	}
	for _, u := range Users {
		if u.Email == a {
			password := []byte(b)
			hashed := []byte(u.Password)
			if err := bcrypt.CompareHashAndPassword(hashed, password); err == nil {
				return true
			}
		}
	}
	return false
}
