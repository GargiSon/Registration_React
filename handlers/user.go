package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"my-react-app/models"
	"my-react-app/mongo"
	"net/http"
	"strconv"
	"time"

	mongodriver "go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(client *mongodriver.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("/api/users called")
		w.Header().Set("Content-Type", "application/json")

		page := 1
		limit := 5
		var err error

		if pageStr := r.URL.Query().Get("page"); pageStr != "" {
			if p, errParse := strconv.Atoi(pageStr); errParse == nil && p > 0 {
				page = p
			}
		}
		if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
			if l, errParse := strconv.Atoi(limitStr); errParse == nil && l > 0 {
				limit = l
			}
		}
		sortField := r.URL.Query().Get("field")
		sortOrder := r.URL.Query().Get("order")

		switch sortField {
		case "username", "email", "mobile":
		case "id", "", "_id":
			sortField = "_id"
		default:
			sortField = "_id"
		}
		if sortOrder != "asc" && sortOrder != "desc" {
			sortOrder = "desc"
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		users, total, err := mongo.GetPaginatedUser(ctx, page, limit, sortField, sortOrder)
		if err != nil {
			log.Printf("Error fetching paginated users: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ApiErrorResponse{Error: "Internal server error"})
			return
		}

		for i := range users {
			if len(users[i].Image) > 0 {
				users[i].ImageBase64 = "data:image/png;base64," + base64.StdEncoding.EncodeToString(users[i].Image)
			} else {
				users[i].ImageBase64 = ""
			}
			users[i].Image = nil
			users[i].Password = ""
		}

		resp := models.UsersListResponse{
			Users:     users,
			Total:     total,
			Page:      page,
			Limit:     limit,
			SortField: sortField,
			SortOrder: sortOrder,
		}
		json.NewEncoder(w).Encode(resp)
	}
}
