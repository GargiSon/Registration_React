package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"my-react-app/mongo"
	"net/http"
	"strconv"
	"time"

	mongodriver "go.mongodb.org/mongo-driver/mongo"
)

// GetUsers returns a handler that fetches sorted, paginated users.
func GetUsers(client *mongodriver.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/api/users called")
		w.Header().Set("Content-Type", "application/json")

		pageStr := r.URL.Query().Get("page")
		limitStr := r.URL.Query().Get("limit")
		sortField := r.URL.Query().Get("field")
		sortOrder := r.URL.Query().Get("order")

		// Safe parsing, use defaults if missing
		page, err := strconv.Atoi(pageStr)
		if err != nil || page <= 0 {
			page = 1
		}
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			limit = 5
		}

		switch sortField {
		case "username", "email", "mobile":
		case "id", "":
			sortField = "_id"
		default:
			sortField = "_id"
		}

		// Only "asc" or "desc" allowed
		if sortOrder != "asc" && sortOrder != "desc" {
			sortOrder = "desc"
		}
		fmt.Printf("Sorting by: %s, Order: %s\n", sortField, sortOrder)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		users, total, err := mongo.GetPaginatedUser(ctx, page, limit, sortField, sortOrder)
		if err != nil {
			fmt.Println("Error fetching paginated users:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		response := map[string]any{
			"users":     users,
			"total":     total,
			"page":      page,
			"limit":     limit,
			"sortField": sortField,
			"sortOrder": sortOrder,
		}
		json.NewEncoder(w).Encode(response)
	}
}
