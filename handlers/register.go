package handlers

import (
	"context"
	"io"
	"my-react-app/models"
	"my-react-app/mongo"
	"net/http"
	"regexp"
	"strings"
	"time"

	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserAPI(client *mongodriver.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse multipart form (max 10MB)
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Could not parse form: "+err.Error(), http.StatusBadRequest)
			return
		}
		// Extract form values
		username := r.FormValue("username")
		password := r.FormValue("password")
		confirm := r.FormValue("confirm")
		email := r.FormValue("email")
		mobile := r.FormValue("mobile")
		address := r.FormValue("address")
		gender := r.FormValue("gender")
		dob := r.FormValue("dob")
		country := r.FormValue("country")
		sports := r.MultipartForm.Value["sports"] // supports multiple checkboxes
		joinedSports := strings.Join(sports, ",")

		if username == "" || password == "" || email == "" || mobile == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}
		if password != confirm {
			http.Error(w, "Passwords do not match", http.StatusBadRequest)
			return
		}
		match, _ := regexp.MatchString(`^(\+\d{1,3})?\d{10}$`, mobile)
		if !match {
			http.Error(w, "Invalid mobile number format", http.StatusBadRequest)
			return
		}
		if dob != "" {
			parsedDOB, err := time.Parse("2006-01-02", dob)
			if err != nil || parsedDOB.After(time.Now()) {
				http.Error(w, "Invalid or future DOB", http.StatusBadRequest)
				return
			}
		}

		// Handle image
		var imageBytes []byte
		file, _, err := r.FormFile("image")
		if err == nil {
			defer file.Close()
			imageBytes, err = io.ReadAll(file)
			if err != nil {
				http.Error(w, "Error reading uploaded image", http.StatusBadRequest)
				return
			}
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Password hashing error", http.StatusInternalServerError)
			return
		}

		// Check email and mobile uniqueness
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if mongo.EmailExists(ctx, email) {
			http.Error(w, "Email already used", http.StatusConflict)
			return
		}
		if mongo.MobileExists(ctx, mobile) {
			http.Error(w, "Mobile already used", http.StatusConflict)
			return
		}

		// Create user struct
		user := models.User{
			Username: username,
			Email:    email,
			Mobile:   mobile,
			Address:  address,
			Gender:   gender,
			Sports:   joinedSports,
			DOB:      dob,
			Country:  country,
			Password: string(hashedPassword),
			Image:    imageBytes,
		}

		// Insert into DB
		err = mongo.InsertUser(ctx, user)
		if err != nil {
			http.Error(w, "Registration failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Success response
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, `{"message":"User registered successfully"}`)
	}
}
