package view

type User struct { //Y
	Id       int
	Email    string
	Username string
	Password string
}
type SessionID struct { //Y
	Email     string
	SessionID string
}
type Post struct { //Y
	Post_ID     int
	User_ID     int
	Post_body   string
	Post_date   string
	Post_time   string
	UserName    string
	Category    string
	Like_counts int

	Comments []Comment
}

type Comment struct {
	Comment_ID   int
	Comment_body string
	User_ID      int
	Post_ID      int
	UserName     string
	Like_counts  int
}

type Like struct {
	User_ID int
	Post_ID int
}
type CommentLike struct {
	Comment_ID int
	User_ID    int
}
 