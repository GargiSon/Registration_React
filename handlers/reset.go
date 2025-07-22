package handlers

import (
	"context"
	"encoding/json"
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
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ResetPasswordResponse{Error: "Invalid or expired token"})
		return
	}

	if time.Now().Unix() > tokenData.TokenExpiry {
		_ = mongo.DeleteResetTokensByUserID(ctx, tokenData.UserID)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ResetPasswordResponse{Error: "Token expired"})
		return
	}

	// Parse body
	var req models.ResetPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.Password != req.Confirm {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ResetPasswordResponse{Error: "Passwords do not match"})
		return
	}

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	err = mongo.UpdateAdminPassword(ctx, tokenData.UserID, string(hashedPass))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ResetPasswordResponse{Error: "Failed to update password"})
		return
	}

	_ = mongo.DeleteResetTokensByUserID(ctx, tokenData.UserID)

	json.NewEncoder(w).Encode(models.ResetPasswordResponse{Message: "Password reset successful"})
}
