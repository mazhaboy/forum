package model

import (
	"fmt"
)

func UpdateSession(NewSessionID, Email string) error {
									  
	_, err := con.Exec("update post set SessionID=? where Email=?", NewSessionID, Email)
	if err != nil {
		return err
	}
	fmt.Println(Email)
	fmt.Println(NewSessionID)
	fmt.Println("ExpiredSessionID is updated")
	return nil
      
}
