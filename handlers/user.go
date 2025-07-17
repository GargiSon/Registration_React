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

func GetUsers(client *mongodriver.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/api/users called")
		w.Header().Set("Content-Type", "application/json")

		pageStr := r.URL.Query().Get("page")
		limitStr := r.URL.Query().Get("limit")

		page, err := strconv.Atoi(pageStr)
		if err != nil || page <= 0 {
			page = 1
		}
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			limit = 5
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		users, total, err := mongo.GetPaginatedUser(ctx, page, limit)
		if err != nil {
			fmt.Println("Error fetching paginated users:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := map[string]any{
			"users": users,
			"total": total,
			"page":  page,
			"limit": limit,
		}

		json.NewEncoder(w).Encode(response)
	}
}
