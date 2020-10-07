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
	Email string
	Post  string
}
