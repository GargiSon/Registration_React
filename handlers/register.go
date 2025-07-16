package handlers

import (
	"context"
	"io"
	"my-react-app/models"
	"my-react-app/mongo"
	"my-react-app/render"
	"my-react-app/utils"
	"net/http"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	countries, err := utils.GetCountriesFromDB()
	if err != nil {
		render.RenderTemplateWithData(w, "Registration.html", models.RegisterPageData{
			Error: "Error fetching countries: " + err.Error(),
		})
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		confirm := r.FormValue("confirm")
		email := r.FormValue("email")
		mobile := r.FormValue("mobile")
		address := r.FormValue("address")
		gender := r.FormValue("gender")
		sports := r.Form["sports"]
		dobStr := r.FormValue("dob")
		country := r.FormValue("country")
		joinedSports := strings.Join(sports, ",")

		user := models.User{
			Username: username,
			Email:    email,
			Mobile:   mobile,
			Address:  address,
			Gender:   gender,
			Sports:   joinedSports,
			DOB:      dobStr,
			Country:  country,
		}

		//sports
		sportsMap := make(map[string]bool)
		for _, s := range sports {
			sportsMap[s] = true
		}

		//password
		if password != confirm {
			render.RenderTemplateWithData(w, "Registration.html", models.RegisterPageData{
				Error:     "Passwords do not match",
				Countries: countries,
				User:      user,
				SportsMap: sportsMap,
			})
			return
		}

		//dob
		dob, err := time.Parse("2006-01-02", dobStr)
		if err != nil || dob.After(time.Now()) {
			render.RenderTemplateWithData(w, "Registration.html", models.RegisterPageData{
				Error:     "Invalid or future DOB",
				Countries: countries,
				User:      user,
				SportsMap: sportsMap,
			})
			return
		}

		//mobile number
		match, err := regexp.MatchString(`^(\+\d{1,3})?\d{10}$`, mobile)
		if err != nil || !match {
			render.RenderTemplateWithData(w, "Registration.html", models.RegisterPageData{
				Error:     "Invalid mobile number format",
				Countries: countries,
				User:      user,
				SportsMap: sportsMap,
			})
			return
		}

		//image
		file, _, err := r.FormFile("image")
		if err == nil {
			defer file.Close()
			imageBytes, err := io.ReadAll(file)
			if err != nil {
				render.RenderTemplateWithData(w, "Registration.html", models.RegisterPageData{
					Error:     "Error in image uploading",
					Countries: countries,
					User:      user,
					SportsMap: sportsMap,
				})
				return
			}
			user.Image = imageBytes //For storing the image
		}

		//hashing password
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			render.RenderTemplateWithData(w, "Registration.html", models.RegisterPageData{
				Error:     "Password hashing failed",
				Countries: countries,
				User:      user,
				SportsMap: sportsMap,
			})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if mongo.EmailExists(ctx, email) {
			render.RenderTemplateWithData(w, "Registration.html", models.RegisterPageData{
				Error:     "Email already used, try a different one.",
				Countries: countries,
				User:      user,
				SportsMap: sportsMap,
			})
			return
		}

		if mongo.MobileExists(ctx, mobile) {
			render.RenderTemplateWithData(w, "Registration.html", models.RegisterPageData{
				Error:     "Mobile number already registered",
				Countries: countries,
				User:      user,
				SportsMap: sportsMap,
			})
			return
		}
		user.Password = string(hashed)

		err = mongo.InsertUser(ctx, user)
		if err != nil {
			render.RenderTemplateWithData(w, "Registration.html", models.RegisterPageData{
				Error:     "Registration failed: " + err.Error(),
				Countries: countries,
				User:      user,
				SportsMap: sportsMap,
			})
			return
		}
		utils.SetFlashMessage(w, "User successfully registered!")
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	render.RenderTemplateWithData(w, "Registration.html", models.RegisterPageData{
		Countries: countries,
		Title:     "Add User",
	})
}
