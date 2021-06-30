package models

type User struct {
	Id      string  `json:"id"`
	Contact Contact `json:"contact"`
}
