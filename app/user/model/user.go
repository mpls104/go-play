package model

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Users []User
