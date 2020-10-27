package model

import (
	"fmt"
)

func UpdateSession(NewSessionID, Email string) error { //Y

	_, err := con.Exec("update Session set Session_ID=? where Email=?", NewSessionID, Email)
	if err != nil {
		return err
	}
	fmt.Println(Email)
	fmt.Println(NewSessionID)
	fmt.Println("ExpiredSessionID is updated")
	return nil

}
func UpdateLike(Likes int, PostID string) error {

	_, err := con.Exec("update pl set Like=? where PostID=?", Likes, PostID)
	if err != nil {
		return err
	}
	fmt.Println(Likes)
	fmt.Println(PostID)
	fmt.Println("Likes is updated")
	return nil

}
