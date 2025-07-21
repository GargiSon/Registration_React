package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username    string             `bson:"username" json:"username"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"-"`
	Mobile      string             `bson:"mobile" json:"mobile"`
	Address     string             `bson:"address" json:"address"`
	Gender      string             `bson:"gender" json:"gender"`
	Sports      string             `bson:"sports" json:"sports"`
	DOB         string             `bson:"dob" json:"dob"`
	Country     string             `bson:"country" json:"country"`
	Image       []byte             `bson:"image,omitempty" json:"-"`
	ImageBase64 string             `json:"imageBase64,omitempty"`
	SNo         int                `bson:"-" json:"sno"`
}

type UsersListResponse struct {
	Users     []User `json:"users"`
	Total     int64  `json:"total"`
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	SortField string `json:"sortField"`
	SortOrder string `json:"sortOrder"`
}

type ApiErrorResponse struct {
	Error string `json:"error"`
}

type EditPageData struct {
	Title     string          `json:"title"`
	User      User            `json:"user"`
	Countries []string        `json:"countries"`
	SportsMap map[string]bool `json:"sportsMap"`
}
