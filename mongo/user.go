package mongo

import (
	"context"
	"my-react-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserCollection() *mongo.Collection {
	return GetCollection("React", "users")
}

func EmailExists(ctx context.Context, email string) bool {
	count, _ := GetUserCollection().CountDocuments(ctx, bson.M{"email": email})
	return count > 0
}

func MobileExists(ctx context.Context, mobile string) bool {
	count, _ := GetUserCollection().CountDocuments(ctx, bson.M{"mobile": mobile})
	return count > 0
}

func InsertUser(ctx context.Context, user models.User) error {
	_, err := GetUserCollection().InsertOne(ctx, user)
	return err
}
