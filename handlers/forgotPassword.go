package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"my-react-app/models"
	"my-react-app/mongo"
	"my-react-app/utils"
	"net/http"
	"os"
	"time"
)

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, models.ForgotResponse{
			Success: false, Message: "Method not allowed",
		})
		return
	}

	var req models.ForgotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Email == "" {
		writeJSON(w, http.StatusBadRequest, models.ForgotResponse{
			Success: false, Message: "Invalid email email.",
		})
		return
	}

	const msg = "If the email exists, a reset link will be sent."
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	admin, err := mongo.GetAdminByEmail(ctx, req.Email)
	if err != nil {
		fmt.Println("Email not found:", req.Email)
		writeJSON(w, http.StatusOK, models.ForgotResponse{Success: true, Message: msg})
		return
	}

	// Token generation
	rawToken := utils.GenerateSecureToken(64)
	tokenHash := utils.HashSHA256(rawToken)
	expiry := time.Now().Add(15 * time.Minute).Unix()

	if err := mongo.InsertResetToken(ctx, admin.ID, tokenHash, expiry); err != nil {
		fmt.Println("DB token insert error:", err)
		writeJSON(w, http.StatusInternalServerError, models.ForgotResponse{
			Success: false, Message: "Server error, please try again later",
		})
		return
	}

	// Reset link
	link := os.Getenv("AUTH_LINK")
	if link == "" {
		link = "http://localhost:5173/reset-password?token="
	}
	resetURL := link + rawToken

	if err := sendResetEmail(req.Email, resetURL); err != nil {
		fmt.Println("Email send error:", err)
		writeJSON(w, http.StatusInternalServerError, models.ForgotResponse{
			Success: false, Message: "Failed to send reset email",
		})
		return
	}

	writeJSON(w, http.StatusOK, models.ForgotResponse{Success: true, Message: msg})
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
