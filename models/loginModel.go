package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
