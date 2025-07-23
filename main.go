package main

import (
	"fmt"
	"log"
	"net/http"

	"my-react-app/handlers"
	"my-react-app/mongo"
	"my-react-app/utils"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	client := mongo.Connect()
	mongo.InitMongoData()

	router := mux.NewRouter()
	utils.SeedDefaultAdmin()

	router.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/api/logout", handlers.LogoutHandler).Methods("GET")
	router.HandleFunc("/api/forgot-password", handlers.ForgotPasswordHandler).Methods("POST")
	router.HandleFunc("/api/reset-password", handlers.ResetHandler).Methods("POST")

	router.HandleFunc("/api/users", handlers.RequireLogin(handlers.GetUsers(client))).Methods("GET")
	router.HandleFunc("/api/users", handlers.RequireLogin(handlers.RegisterUserAPI(client))).Methods("POST")
	router.HandleFunc("/api/users/{id}", handlers.RequireLogin(handlers.DeleteHandler(client))).Methods("DELETE")
	router.HandleFunc("/api/users/{id}", handlers.RequireLogin(handlers.GetUserHandler(client))).Methods("GET")
	router.HandleFunc("/api/users/{id}", handlers.RequireLogin(handlers.UpdateHandler(client))).Methods("PUT")

	router.HandleFunc("/api/countries", handlers.RequireLogin(handlers.GetCountries)).Methods("GET")

	handler := cors.AllowAll().Handler(router)

	fmt.Println("Server running at http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", handler))
}
