package handlers

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func sendResetEmail(toEmail, resetLink string) error {
	email := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	if email == "" || password == "" {
		return fmt.Errorf("SMTP credentials not found in environment")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Password Reset Link")
	m.SetBody("text/html", fmt.Sprintf(`Hi,
	We received a request to reset your password. Click the link below to set a new password: %s
	If you didn't request this, you can safely ignore this email.
	Thanks,
	Gargi`, resetLink))

	d := gomail.NewDialer("smtp.gmail.com", 587, email, password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send mail: %w", err)
	}
	return nil
}
