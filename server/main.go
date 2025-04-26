package main

import (
	"cosmic-gate-chat/config"
	"cosmic-gate-chat/handlers"
	"cosmic-gate-chat/websocket"
	"log"
	"net/http"

	serverHandlers "github.com/gorilla/handlers"
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
	origins := serverHandlers.AllowedOrigins([]string{"*"})
	methods := serverHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	headers := serverHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Routes
	router.HandleFunc("/auth", handlers.AuthUserHandler).Methods("POST")

	router.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", handlers.UpdateUserHandler).Methods("PATCH")
	router.HandleFunc("/users/all/{userId}", handlers.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/users/friends", handlers.GetUserFriendsHandler).Methods("GET")
	router.HandleFunc("/users/friend-requests", handlers.SendFriendRequestHandler).Methods("POST")
	router.HandleFunc("/users/friend-requests", handlers.GetUserFriendRequestsHandler).Methods("GET")

	router.HandleFunc("/friend-request/accept", handlers.AcceptFriendRequestHandler).Methods("POST")
	router.HandleFunc("/friend-request/decline", handlers.DeclineFriendRequestHandler).Methods("POST")

	router.HandleFunc("/messages", handlers.GetMessagesHandler).Methods("GET")

	// WebSocket
	router.HandleFunc("/ws/{senderId}", websocket.HandleWebSocket)

	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", serverHandlers.CORS(origins, methods, headers)(router)))
}
