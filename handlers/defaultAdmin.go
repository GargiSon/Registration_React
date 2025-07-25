package handlers

import (
	"encoding/json"
	"my-react-app/utils"
	"net/http"
)

func SeedAdminHandler(w http.ResponseWriter, r *http.Request) {
	utils.SeedDefaultAdmin()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Seed function called. Check logs for output.",
	})
}
