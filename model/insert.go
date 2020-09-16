package model

import (
	"fmt"
	"log"
)

func Insert(a, b, c string) error {

	_, err := con.Exec("INSERT INTO users (username,email,password) VALUES(?,?,?)", a, b, c)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Data is inserted")
	return nil
	
}
