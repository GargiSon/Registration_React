package handlers

import (
	"encoding/json"
	"my-react-app/utils"
	"net/http"
)

func GetCountries(w http.ResponseWriter, r *http.Request) {
	countries, err := utils.GetCountriesFromDB()
	if err != nil {
		http.Error(w, "Failed to fetch countries", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}
