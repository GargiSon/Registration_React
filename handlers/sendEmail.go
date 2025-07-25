package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

func sendResetEmail(toEmail, resetLink string) error {
	tmpl, err := template.ParseFiles("templates/sendEmail.html")
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	var bodyBuffer bytes.Buffer
	err = tmpl.Execute(&bodyBuffer, struct{ Link string }{Link: resetLink})
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}
	email := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	if email == "" || password == "" {
		return fmt.Errorf("SMTP credentials not found in environment")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Password Reset Link")
	m.SetBody("text/html", bodyBuffer.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, email, password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send mail: %w", err)
	}
	return nil
}
