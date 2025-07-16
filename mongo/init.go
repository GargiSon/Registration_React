package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func InitMongoData() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := os.Getenv("MONGO_DB_NAME")
	if db == "" {
		log.Println("MONGO_DB_NAME not set")
		return
	}

	countryColl := GetCollection(db, "countries")
	countryCount, err := countryColl.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Println("Error checking countries:", err)
	} else if countryCount == 0 {
		countries := []any{
			bson.M{"name": "INDIA"},
			bson.M{"name": "AFGHANISTHAN"},
			bson.M{"name": "FRANCE"},
		}
		if _, err := countryColl.InsertMany(ctx, countries); err != nil {
			log.Println("Failed to insert default countries:", err)
		} else {
			fmt.Println("Inserted default countries.")
		}
	} else {
		fmt.Println("Countries already exist.")
	}
}
