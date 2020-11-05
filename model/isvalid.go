package model

import (
	"fmt"
	"log"

	"github.com/mazhaboy/forum/tree/master/view"
	"golang.org/x/crypto/bcrypt"
	// view "../view"
)

func IsValid(a, b string) bool { //Y

	rows, err := con.Query("select * from User")
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

func IsUserValid(Session string) bool { //Y

	rows, err := con.Query("select * from Session")
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
func GetPosts(filter string, User_ID int) []view.Post {
	var post string

	rowss, errr := con.Query("select t1.*, count(t2.User_ID) as LikeCount from Comment t1 left join LikeComment t2 USING(Comment_ID) group by t1.Comment_ID")
	if errr != nil {
		log.Fatal(errr)
	}
	Comments := []view.Comment{}
	for rowss.Next() {
		p := view.Comment{}
		err := rowss.Scan(&p.Comment_ID, &p.Comment_body, &p.User_ID, &p.Post_ID, &p.UserName, &p.Like_counts)

		if err != nil {
			fmt.Println("Error")
			continue
		}
		Comments = append(Comments, p)
	}
	fmt.Println("Rabotaet")

	if filter == "sport" {
		post = "select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t1.Categories='sport' group by t1.Post_ID order by Post_ID DESC"
	} else if filter == "religion" {
		post = "select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t1.Categories='religion' group by t1.Post_ID order by Post_ID DESC"
	} else if filter == "politics" {
		post = "select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t1.Categories='politics' group by t1.Post_ID order by Post_ID DESC"
	} else if filter == "science" {
		post = "select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t1.Categories='science' group by t1.Post_ID order by Post_ID DESC"
	} else if filter == "others" {
		post = "select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t1.Categories='others' group by t1.Post_ID order by Post_ID DESC"
	} else {
		post = "select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) group by t1.Post_ID order by Post_ID DESC"
	}
	if filter == "myposts" {
		fmt.Println("aksjfkjashfkjashkjhakjsfhjkashfhafakjsfhkajshfkjashf,", User_ID)
		rows, err := con.Query("select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t1.User_ID=? group by t1.Post_ID order by Post_ID DESC", User_ID)
		if err != nil {
			log.Fatal(err)
		}
		Posters := []view.Post{}
		for rows.Next() {
			p := view.Post{}
			err := rows.Scan(&p.Post_ID, &p.User_ID, &p.Post_body, &p.Post_date, &p.Post_time, &p.UserName, &p.Category, &p.Like_counts)

			if err != nil {
				fmt.Println("Error")
				continue
			}
			Posters = append(Posters, p)
		}
		fmt.Println("Rabotaet")
		i := 0
		for _, p := range Posters {
			for _, c := range Comments {
				if p.Post_ID == c.Post_ID {
					p.Comments = append(p.Comments, c)
				}

			}
			Posters[i].Comments = p.Comments
			i++

		}

		return Posters
	}
	if filter == "myfavourite" {
		fmt.Println("kkkkkkkkkkkk", User_ID)
		rows, err := con.Query("select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t2.User_ID=? group by t1.Post_ID order by Post_ID DESC", User_ID)
		if err != nil {
			log.Fatal(err)
		}
		Posters := []view.Post{}
		for rows.Next() {
			p := view.Post{}
			err := rows.Scan(&p.Post_ID, &p.User_ID, &p.Post_body, &p.Post_date, &p.Post_time, &p.UserName, &p.Category, &p.Like_counts)

			if err != nil {
				fmt.Println("Error")
				continue
			}
			Posters = append(Posters, p)
		}
		fmt.Println("Rabotaet")
		i := 0
		for _, p := range Posters {
			for _, c := range Comments {
				if p.Post_ID == c.Post_ID {
					p.Comments = append(p.Comments, c)
				}

			}
			Posters[i].Comments = p.Comments
			i++

		}

		return Posters
	}
	rows, err := con.Query(post)
	if err != nil {
		log.Fatal(err)
	}
	Posters := []view.Post{}
	for rows.Next() {
		p := view.Post{}
		err := rows.Scan(&p.Post_ID, &p.User_ID, &p.Post_body, &p.Post_date, &p.Post_time, &p.UserName, &p.Category, &p.Like_counts)

		if err != nil {
			fmt.Println("Error")
			continue
		}
		Posters = append(Posters, p)
	}
	fmt.Println("Rabotaet")
	i := 0
	for _, p := range Posters {
		for _, c := range Comments {
			if p.Post_ID == c.Post_ID {
				p.Comments = append(p.Comments, c)
			}

		}
		Posters[i].Comments = p.Comments
		i++

	}
	fmt.Println("jsfjkasfkjashfkjashfjsahfkjasfhkjashfjafkjaaaaaaaaaaaaa")
	return Posters

}
func GetUserIDbySession(ID string) (int, string) {
	UserName := "..."
	User_ID := 0
	rows, err := con.Query("select * from Session")
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
		if s.SessionID == ID {
			User_ID, UserName = GetUsrIDandName(s.Email)
			fmt.Println("Success1")

		}
	}
	return User_ID, UserName

}
func GetUsrIDandName(Email string) (int, string) {

	UserName := ""
	User_ID := 0
	rows, err := con.Query("select * from User")
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
		if u.Email == Email {
			User_ID = u.Id
			UserName = u.Username
			fmt.Println("Success2")
		}
	}
	return User_ID, UserName

}

func IsLiked(User_ID int, Post_ID int) bool {

	rows, err := con.Query("select * from Like")
	if err != nil {
		log.Fatal(err)
	}

	Users := []view.Like{}

	for rows.Next() {
		s := view.Like{}
		err := rows.Scan(&s.User_ID, &s.Post_ID)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		Users = append(Users, s)

	}

	for _, s := range Users {
		if s.User_ID == User_ID && s.Post_ID == Post_ID {
			_, err := con.Exec("delete from Like where User_ID=? and Post_ID=?", User_ID, Post_ID)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Like is deleted")
			return true
		}
	}
	return false
}
func IsCommentLiked(Comment_ID int, User_ID int) bool {

	rows, err := con.Query("select * from LikeComment")
	if err != nil {
		log.Fatal(err)
	}

	Users := []view.CommentLike{}

	for rows.Next() {
		s := view.CommentLike{}
		err := rows.Scan(&s.Comment_ID, &s.User_ID)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		Users = append(Users, s)

	}

	for _, s := range Users {
		if s.User_ID == User_ID && s.Comment_ID == Comment_ID {
			_, err := con.Exec("delete from LikeComment where User_ID=? and Comment_ID=?", User_ID, Comment_ID)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("CommentLike is deleted")
			return true
		}
	}
	return false
}

// func AddComment(User_ID int, Post_ID int) bool {

// 	rows, err := con.Query("select * from Like")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	Users := []view.Like{}

// 	for rows.Next() {
// 		s := view.Like{}
// 		err := rows.Scan(&s.User_ID, &s.Post_ID)
// 		if err != nil {
// 			fmt.Println("Error")
// 			continue
// 		}
// 		Users = append(Users, s)

// 	}

// 	for _, s := range Users {
// 		if s.User_ID == User_ID && s.Post_ID == Post_ID {
// 			_, err := con.Exec("delete from Like where User_ID=? and Post_ID=?", User_ID, Post_ID)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			fmt.Println("Like is deleted")
// 			return true
// 		}
// 	}
// 	return false
// }
