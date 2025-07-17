package mongo

import (
	"context"
	"my-react-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPaginatedUser(ctx context.Context, page, limit int) ([]models.User, int64, error) {
	offset := (page - 1) * limit

	findOptions := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(limit))

	cursor, err := GetUserCollection().Find(ctx, bson.M{}, findOptions)

	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close(ctx)

	var users []models.User

	if err := cursor.All(ctx, &users); err != nil {
		return nil, 0, err
	}

	total, err := GetUserCollection().CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}
