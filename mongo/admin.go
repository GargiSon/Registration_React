package mongo

import (
	"context"
	"log"
	"my-react-app/models"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func getDBName() string {
	return os.Getenv("MONGO_DB_NAME")
}

func GetAdminByEmail(ctx context.Context, email string) (models.Admin, error) {
	var admin models.Admin
	collection := GetCollection(getDBName(), "admins")
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&admin)
	return admin, err
}

func CheckAdminExists(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := GetCollection(getDBName(), "admins")
	count, err := collection.CountDocuments(ctx, bson.M{"email": email})
	return count > 0, err
}

func UpdateAdminByEmail(ctx context.Context, email string, hashedPassword string) error {
	result, err := GetCollection(getDBName(), "admins").
		UpdateOne(ctx, bson.M{"email": email}, bson.M{
			"$set": bson.M{"password": hashedPassword},
		})

	if err != nil {
		log.Println("Update error:", err)
		return err
	}

	log.Printf("Matched %v document(s), Modified %v document(s)\n", result.MatchedCount, result.ModifiedCount)
	return nil
}
