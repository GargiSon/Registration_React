package mongo

import (
	"context"
	"fmt"
	"log"
	"my-react-app/models"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
)

func getDBName() string {
	return os.Getenv("MONGO_DB_NAME")
}

func GetAdminByID(ctx context.Context, userID primitive.ObjectID, admin *models.Admin) error {
	collection := GetCollection(getDBName(), "admins")
	return collection.FindOne(ctx, bson.M{"_id": userID}).Decode(admin)
}

func GetAdminByEmail(ctx context.Context, email string) (models.Admin, error) {
	var admin models.Admin
	collection := GetCollection(getDBName(), "admins")
	email = strings.ToLower(email)
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
	email = strings.ToLower(email)
	collection := GetCollection(getDBName(), "admins")

	// Check if admin exists first
	var admin models.Admin
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&admin)
	if err == mongodriver.ErrNoDocuments {
		log.Printf("No admin found with email: %s\n", email)
		return fmt.Errorf("admin not found")
	} else if err != nil {
		log.Printf("FindOne error: %v\n", err)
		return err
	}

	// Proceed to update password
	result, err := collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{
		"$set": bson.M{"password": hashedPassword},
	})
	if err != nil {
		log.Println("Update error:", err)
		return err
	}

	log.Printf("Matched %v document(s), Modified %v document(s)\n", result.MatchedCount, result.ModifiedCount)
	if result.MatchedCount == 0 {
		return fmt.Errorf("no admin matched the email: %s", email)
	}
	return nil
}
