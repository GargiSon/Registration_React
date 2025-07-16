package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"my-react-app/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/api/users called")
		w.Header().Set("Content-Type", "application/json")
		collection := client.Database("React").Collection("user")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			fmt.Println("Error finding users:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)

		var users []models.User
		if err := cursor.All(ctx, &users); err != nil {
			fmt.Println("Cursor decode error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("Users fetched:", users)
		json.NewEncoder(w).Encode(users)
	}
}
