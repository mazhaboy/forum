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
