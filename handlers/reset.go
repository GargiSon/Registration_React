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

	// Only POST supported
	if r.Method != http.MethodPost {
		log.Println("ResetHandler: Method not allowed")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	rawToken := r.URL.Query().Get("token")
	if rawToken == "" {
		log.Println("ResetHandler: No token provided")
		http.Error(w, "Token not provided", http.StatusBadRequest)
		return
	}

	tokenHash := utils.HashSHA256(rawToken)
	log.Println("ResetHandler: Token hash generated")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tokenData, err := mongo.FindResetToken(ctx, tokenHash)
	if err != nil {
		log.Println("ResetHandler: Invalid or expired token:", err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ResetPasswordResponse{Error: "Invalid or expired token"})
		return
	}
	log.Printf("ResetHandler: Found token for user ID %s\n", tokenData.UserID.Hex())

	if time.Now().Unix() > tokenData.TokenExpiry {
		log.Println("ResetHandler: Token expired, deleting from DB")
		_ = mongo.DeleteResetTokensByUserID(ctx, tokenData.UserID)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ResetPasswordResponse{Error: "Token expired"})
		return
	}

	// Parse body
	var req models.ResetPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("ResetHandler: Error decoding JSON body:", err)
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	log.Println("ResetHandler: Parsed JSON body")

	if req.Password != req.Confirm {
		log.Println("ResetHandler: Passwords do not match")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ResetPasswordResponse{Error: "Passwords do not match"})
		return
	}
	log.Println("ResetHandler: Passwords match")

	// Hash the new password securely using bcrypt
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("ResetHandler: Failed to hash password:", err)
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	log.Println("ResetHandler: Password hashed successfully")

	err = mongo.UpdateAdminByEmail(ctx, tokenData.Email, string(hashedPass))

	if err != nil {
		log.Println("ResetHandler: Failed to update password in DB:", err)
		http.Error(w, "Failed to update password", http.StatusInternalServerError)
		return
	}
	log.Printf("ResetHandler: Password updated for user ID %s\n", tokenData.UserID.Hex())

	err = mongo.DeleteResetTokensByUserID(ctx, tokenData.UserID)
	if err != nil {
		log.Println("ResetHandler: Failed to delete reset token after successful reset:", err)
	} else {
		log.Println("ResetHandler: Reset token deleted successfully")
	}

	json.NewEncoder(w).Encode(models.ResetPasswordResponse{Message: "Password reset successful"})
	log.Println("ResetHandler: Password reset process completed")
}
