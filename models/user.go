package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
}
