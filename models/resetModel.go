package models

type ResetPasswordRequest struct {
	Password string `json:"password"`
	Confirm  string `json:"confirm"`
}

type ResetPasswordResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
