package main

import (
	"cosmic-gate-chat/config"
	"cosmic-gate-chat/controllers"
	"cosmic-gate-chat/websocket"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	// Connect to the DB
	config.InitMongoDB()

	// Routes Setting
	router := mux.NewRouter()
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Routes
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetUser).Methods("GET")

	router.HandleFunc("/messages", controllers.GetMessages).Methods("GET")

	// WebSocket
	router.HandleFunc("/ws/{senderId}", websocket.HandleWebSocket)

	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(origins, methods, headers)(router)))
}
