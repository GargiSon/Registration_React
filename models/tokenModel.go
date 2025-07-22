package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PasswordResetToken struct {
	UserID      primitive.ObjectID `bson:"user_id"`
	TokenHash   string             `bson:"token"`
	TokenExpiry int64              `bson:"token_expiry"`
}
