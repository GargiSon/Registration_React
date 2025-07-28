package handlers

import (
	"context"
	"encoding/json"
	"log"
	"my-react-app/models"
	"my-react-app/mongo"
	"my-react-app/utils"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	setNoCacheHeaders(w)

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	rawToken := r.URL.Query().Get("token")
	if rawToken == "" {
		http.Error(w, "Token not provided", http.StatusBadRequest)
		return
	}

	tokenHash := utils.HashSHA256(rawToken)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tokenData, err := mongo.FindResetToken(ctx, tokenHash)
	if err != nil {
		writeJSONError(w, http.StatusUnauthorized, "Invalid or expired token")
		return
	}

	if time.Now().Unix() > tokenData.TokenExpiry {
		if delErr := mongo.DeleteResetTokensByUserID(ctx, tokenData.UserID); delErr != nil {
			log.Println("Failed to delete expired token:", delErr)
		}
		writeJSONError(w, http.StatusUnauthorized, "Token expired")
		return
	}

	var admin models.Admin
	if err := mongo.GetAdminByID(ctx, tokenData.UserID, &admin); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to find admin for token")
		return
	}
	if admin.Email == "" {
		writeJSONError(w, http.StatusInternalServerError, "Admin email missing")
		return
	}

	defer r.Body.Close()
	var req models.ResetPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if req.Password != req.Confirm {
		writeJSONError(w, http.StatusBadRequest, "Passwords do not match")
		return
	}

	if len(req.Password) > 128 {
		writeJSONError(w, http.StatusBadRequest, "Password too long")
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	err = mongo.UpdateAdminByEmail(ctx, admin.Email, string(hashedPass))
	if err != nil {
		http.Error(w, "Failed to update password", http.StatusInternalServerError)
		return
	}

	if err = mongo.DeleteResetTokensByUserID(ctx, tokenData.UserID); err != nil {
		log.Println("Failed to delete reset token after successful reset:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ResetPasswordResponse{Message: "Password reset successful"})
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(models.ResetPasswordResponse{Error: message})
}
