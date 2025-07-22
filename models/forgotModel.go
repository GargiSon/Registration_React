package models

type ForgotRequest struct {
	Email string `json:"email"`
}

type ForgotResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
