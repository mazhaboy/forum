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
func IsUserValid(Session string) bool {

	rows, err := con.Query("select * from post")
	if err != nil {
		log.Fatal(err)
	}
	Users := []view.SessionID{}
	for rows.Next() {
		s := view.SessionID{}
		err := rows.Scan(&s.Email, &s.SessionID)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		Users = append(Users, s)
	}
	for _, s := range Users {
		if s.SessionID == Session {
			fmt.Println("Checked")
			return true

		}
	}
	return false

}
func GetPosts() []view.Posts {

	rows, err := con.Query("select * from pl")
	if err != nil {
		log.Fatal(err)
	}
	Posters := []view.Posts{}
	for rows.Next() {
		p := view.Posts{}
		err := rows.Scan(&p.Email, &p.Post)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		Posters = append(Posters, p)
	}
	return Posters

}
