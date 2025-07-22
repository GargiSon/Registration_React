package main

import (
	"fmt"
	"log"
	"net/http"

	"my-react-app/handlers"
	"my-react-app/mongo"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	client := mongo.Connect()
	mongo.InitMongoData()

	router := mux.NewRouter()
	router.HandleFunc("/api/users", handlers.GetUsers(client)).Methods("GET")
	router.HandleFunc("/api/users", handlers.RegisterUserAPI(client)).Methods("POST")
	router.HandleFunc("/api/users/{id}", handlers.DeleteHandler(client)).Methods("DELETE")
	router.HandleFunc("/api/users/{id}", handlers.GetUserHandler(client)).Methods("GET")
	router.HandleFunc("/api/users/{id}", handlers.UpdateHandler(client)).Methods("PUT")

	router.HandleFunc("/api/countries", handlers.GetCountries).Methods("GET")
	router.HandleFunc("/api/forgot-password", handlers.ForgotPasswordHandler).Methods("POST")
	router.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/api/logout", handlers.LogoutHandler).Methods("GET")
	router.HandleFunc("/api/reset-password", handlers.ResetHandler).Methods("POST")

	handler := cors.AllowAll().Handler(router)

	fmt.Println("Server running at http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", handler))
}
