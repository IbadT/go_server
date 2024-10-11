package models

type User struct {
	ID       string `json:"ID"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
