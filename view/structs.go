package view

type User struct {
	Id       int
	Email    string
	Username string
	Password string
}
type SessionID struct {
	Email     string
	SessionID string
}
type Posts struct {
	PostID int
	Email  string
	Post   string
	Like   int
}
