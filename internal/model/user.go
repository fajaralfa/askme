package model

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Photo    string `json:"photo"`
}
