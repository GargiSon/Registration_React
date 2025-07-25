package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	sessionStore  = make(map[string]string)
	sessionMutex  = sync.Mutex{}
	userPageLimit int
)

func InitSession() {
	if limitStr := os.Getenv("USER_PAGE_LIMIT"); limitStr != "" {
		if val, err := strconv.Atoi(limitStr); err == nil && val > 0 {
			userPageLimit = val
			return
		}
	}
	userPageLimit = 5
}

func GenerateSecureToken(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func SetSession(w http.ResponseWriter, email string) {
	token := GenerateSecureToken(64)

	sessionMutex.Lock()
	sessionStore[token] = email
	sessionMutex.Unlock()

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
}

func GetSessionEmail(r *http.Request) (string, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie.Value == "" {
		return "", false
	}

	sessionMutex.Lock()
	defer sessionMutex.Unlock()

	email, ok := sessionStore[cookie.Value]
	return email, ok
}

func ClearSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		sessionMutex.Lock()
		delete(sessionStore, cookie.Value)
		sessionMutex.Unlock()

		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteNoneMode,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Logged out successfully",
	})
}

func GetUserPageLimit() int {
	return userPageLimit
}

func RequireLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setNoCacheHeaders(w)

		_, ok := GetSessionEmail(r)
		if !ok {
			accept := r.Header.Get("Accept")
			if strings.Contains(accept, "application/json") || strings.HasPrefix(r.URL.Path, "/api/") {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
				return
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func setNoCacheHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}
