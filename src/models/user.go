package models

type User struct {
	Id   string `sql:"type:uuid;primary_key;"`
	Name string `json:"name"`
}
