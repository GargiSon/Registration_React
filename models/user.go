package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username    string             `bson:"username" json:"username"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password"`
	Mobile      string             `bson:"mobile" json:"mobile"`
	Address     string             `bson:"address" json:"address"`
	Gender      string             `bson:"gender" json:"gender"`
	Sports      string             `bson:"sports" json:"sports"`
	DOB         string             `bson:"dob" json:"dob"`
	Country     string             `bson:"country" json:"country"`
	Image       []byte             `bson:"image,omitempty" json:"image,omitempty"`
	ImageBase64 string             `json:"imageBase64"`
}

type RegisterPageData struct {
	User      User
	Countries []string
	SportsMap map[string]bool
	Error     string
	Title     string
}
