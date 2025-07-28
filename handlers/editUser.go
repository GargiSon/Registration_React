package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"my-react-app/models"
	"my-react-app/mongo"
	"my-react-app/utils"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
)

func GetUserHandler(client *mongodriver.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		idStr := vars["id"]
		if idStr == "" {
			http.Error(w, `{"error":"Missing user ID"}`, http.StatusBadRequest)
			return
		}

		objID, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			http.Error(w, `{"error":"Invalid user ID"}`, http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		user, err := mongo.FindUserByID(ctx, objID)
		if err != nil {
			http.Error(w, `{"error":"User not found"}`, http.StatusNotFound)
			return
		}

		if len(user.Image) > 0 {
			user.ImageBase64 = base64.StdEncoding.EncodeToString(user.Image)
		}
		if len(user.DOB) > 10 {
			user.DOB = user.DOB[:10]
		}

		countries, _ := utils.GetCountriesFromDB()

		sportsMap := make(map[string]bool)
		for _, sport := range strings.Split(user.Sports, ",") {
			s := strings.TrimSpace(sport)
			if s != "" {
				sportsMap[s] = true
			}
		}

		response := models.EditPageData{
			Title:     "Edit User",
			User:      user,
			Countries: countries,
			SportsMap: sportsMap,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func UpdateHandler(client *mongodriver.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost && r.Method != http.MethodPut {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		vars := mux.Vars(r)
		idStr := vars["id"]
		if idStr == "" {
			http.Error(w, `{"error":"Missing user ID in URL"}`, http.StatusBadRequest)
			return
		}

		objID, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			http.Error(w, `{"error":"Invalid user ID"}`, http.StatusBadRequest)
			return
		}

		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, `{"error":"Failed to parse form"}`, http.StatusBadRequest)
			return
		}

		username := r.FormValue("username")
		mobile := r.FormValue("mobile")
		address := r.FormValue("address")
		gender := r.FormValue("gender")
		dob := r.FormValue("dob")
		country := r.FormValue("country")
		sports := strings.Join(r.Form["sports"], ",")
		removeImage := r.FormValue("remove_image") == "1"

		match, _ := regexp.MatchString(`^(\+\d{1,3})?\d{10}$`, mobile)
		if !match {
			http.Error(w, `{"error":"Invalid mobile format"}`, http.StatusBadRequest)
			return
		}

		if dob != "" {
			parsedDOB, err := time.Parse("2006-01-02", dob)
			if err != nil || parsedDOB.After(time.Now()) {
				http.Error(w, "Invalid or future DOB", http.StatusBadRequest)
				return
			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		update := bson.M{
			"username": username,
			"mobile":   mobile,
			"address":  address,
			"gender":   gender,
			"sports":   sports,
			"dob":      dob,
			"country":  country,
		}

		// Handle image: update if uploaded, remove if requested, else preserve
		file, _, err := r.FormFile("image")
		if err == nil {
			defer file.Close()
			imageData, _ := io.ReadAll(file)
			update["image"] = imageData
		} else if removeImage {
			update["image"] = nil
		} else {
			// Do not update image field — preserve existing
		}

		err = mongo.UpdateUserByID(ctx, objID, update)
		if err != nil {
			utils.SetFlashMessage(w, "Update failed: "+err.Error())
			http.Error(w, `{"error":"Update failed"}`, http.StatusInternalServerError)
			return
		}

		utils.SetFlashMessage(w, "User successfully updated!")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User successfully updated!",
		})
	}
}
