package model

import (
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
)

func UpdateSession(ExpiredSessionID string) error {
	u2, _ := uuid.NewV4()
	NewSessionID := u2.String()
	_, err := con.Exec("update post set SessionID=? where SessionID=?", NewSessionID, ExpiredSessionID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ExpiredSessionID is updated")
	return nil

}
