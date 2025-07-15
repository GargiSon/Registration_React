package models

type User struct {
	ID     string `json:"id" bson:"_id, omitempty"`
	Name   string `json:"name" bson:"name"`
	Email  string `json:"email" bson:"email"`
	Mobile string `json:"mobile" bson:"mobile"`
}
