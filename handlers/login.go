package handlers

import (
	"context"
	"encoding/json"
	"log"
	"my-react-app/models"
	"my-react-app/mongo"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	setNoCacheHeaders(w)

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decoding login request:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password must be provided", http.StatusBadRequest)
		return
	}

	// Context for DB operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Fetch user
	admin, err := mongo.GetAdminByEmail(ctx, req.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.LoginResponse{Error: "Invalid email or password"})
		return
	}

	log.Printf("DB hash for %s: %s", req.Email, admin.Password)
	log.Printf("Entered password: %s", req.Password)

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		// Log for debugging but don't reveal to client
		log.Printf("Password mismatch for user: %s", req.Email)
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Set session (cookie + in-memory map)
	SetSession(w, req.Email)

	// Success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.LoginResponse{
		Message: "Login Successful!",
	})
}
