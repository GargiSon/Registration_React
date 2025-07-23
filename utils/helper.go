package utils

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"my-react-app/mongo"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func GetCountriesFromDB() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongo.GetCollection("React", "countries")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var countries []string
	for cursor.Next(ctx) {
		var doc struct {
			Name string `bson:"name"`
		}
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		countries = append(countries, doc.Name)
	}
	return countries, nil
}

func SeedDefaultAdmin() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	email := "gargi.soni@loginradius.com"
	plainPassword := "tyu"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("❌ Error creating password hash:", err)
		return
	}
	collection := mongo.GetCollection("React", "admins")

	filter := bson.M{"email": email}
	update := bson.M{
		"$set": bson.M{
			"email":    email,
			"password": string(hashedPassword),
		},
	}

	opts := options.Update().SetUpsert(true)

	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println("❌ Failed inserting/updating admin user:", err)
		return
	}
	log.Printf("✅ Admin seeded. Matched: %v, Updated: %v, Upserted: %v",
		result.MatchedCount, result.ModifiedCount, result.UpsertedID)
}

func SetFlashMessage(w http.ResponseWriter, message string) {
	http.SetCookie(w, &http.Cookie{
		Name:  "flash",
		Value: message,
		Path:  "/",
	})
}

func GetFlashMessage(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("flash")
	if err != nil {
		return ""
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "flash",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	return cookie.Value
}

func GenerateSecureToken(length int) string {
	bytes := make([]byte, length)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func HashSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
