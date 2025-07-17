package mongo

import (
	"context"
	"my-react-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPaginatedUser(ctx context.Context, page, limit int, sortField, sortOrder string) ([]models.User, int64, error) {
	offset := (page - 1) * limit

	if sortField == "sno" {
		sortField = "_id"
	}

	// Map "sno" and "id" to MongoDB's internal "_id".
	if sortField == "sno" || sortField == "id" || sortField == "" {
		sortField = "_id"
	}

	// Supported user-sortable fields
	switch sortField {
	case "username", "email", "mobile", "_id":
		// Ok
	default:
		sortField = "_id"
	}

	// Allowed order "asc" or "desc"
	sortDirection := getSortOrderValue(sortOrder)

	findOptions := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: sortField, Value: sortDirection}})

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

	// Set SNo field for frontend display if needed
	startSNo := (page-1)*limit + 1
	for i := range users {
		users[i].SNo = startSNo + i
	}
	return users, total, nil
}

func getSortOrderValue(order string) int {
	if order == "asc" {
		return 1
	}
	return -1
}
