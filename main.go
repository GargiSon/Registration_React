package main

import (
	"fmt"
	"log"
	"net/http"

	"my-react-app/handlers"
	db "my-react-app/mongo"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	client := db.ConnectDB()

	router := mux.NewRouter()
	router.HandleFunc("/api/users", handlers.GetUsers(client)).Methods("GET")

	handler := cors.AllowAll().Handler(router)

	fmt.Println("Server running at http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", handler))
}
