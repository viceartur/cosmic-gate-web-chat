package websocket

import (
	"cosmic-gate-chat/models"
	"cosmic-gate-chat/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type MessageWS struct {
	Type        string `json:"type"`
	SenderID    string `json:"senderId,omitempty"`
	RecipientID string `json:"recipientId,omitempty"`
	Data        string `json:"data"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var clients = make(map[string]*websocket.Conn)
var clientsMutex sync.Mutex

// The method handles WebSocket Requests
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Update Upgrader options
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	// Take out the Sender ID from parameters
	vars := mux.Vars(r)
	senderId := vars["senderId"]

	// Add connected Sender ID as a Client
	addClient(senderId, ws)

	// Read Incoming Messages
	go reader(senderId, ws)
}

// Store a Client Connection
func addClient(clientId string, conn *websocket.Conn) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	clients[clientId] = conn
	log.Println(clientId, "connected")
}

// Handle Incoming WS Messages
func reader(userId string, conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("ReadMessage error:", err)
			// Check whether the User Connection exists
			if _, ok := clients[userId]; ok {
				clients[userId].Close() // close connection
				delete(clients, userId) // remove connection
				log.Println("Client closed and removed: ", userId)
			}
			return
		}

		jsonData := string(p)
		var message MessageWS
		err = json.Unmarshal([]byte(jsonData), &message)
		if err != nil {
			log.Println("WS Reader error parsing JSON: ", err)
		}

		// Based on Message Type handle the Message
		if message.Type == "chat-connection" {
			handleChatConnection(userId, message)
		} else if message.Type == "chat-message" {
			handleChatMessage(userId, message)
		} else if message.Type == "friend-request-sent" {
			handleFriendRequestSent(userId, message)
		}
	}
}

// Notify the Client that Recipient connected to the chat
func handleChatConnection(userId string, message MessageWS) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	// Convert Notification into JSON
	messageToSend := MessageWS{
		Type:        message.Type,
		SenderID:    userId,
		RecipientID: message.RecipientID,
		Data:        "joined to the chat",
	}
	jsonMessage, err := json.Marshal(messageToSend)
	if err != nil {
		log.Println(err)
	}

	// Find a Recipient in the Connections
	client, ok := clients[message.RecipientID]
	if !ok {
		log.Println("Recipient not connected yet. ID:", message.RecipientID)
		return
	}

	// Send Message to the Recipient
	err = client.WriteMessage(websocket.TextMessage, jsonMessage)
	if err != nil {
		log.Println("WS Broadcast WriteMessage error:", err)
	}
}

// Send Message from the Client to Recipient
func handleChatMessage(userId string, message MessageWS) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	jsonMsg, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	// Save Message into DB
	messageToSave := models.Message{
		SenderID:    userId,
		RecipientID: message.RecipientID,
		Text:        message.Data,
		SentAt:      time.Now(),
	}
	go services.SaveMessage(messageToSave)

	// Find a Recipient in the Connections
	client, ok := clients[message.RecipientID]
	if !ok {
		log.Println("Recipient not connected yet. ID:", message.RecipientID)
		return
	}

	// Send Message to the Recipient
	err = client.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		log.Println("WS Broadcast WriteMessage error:", err)

	}
}

// Send a Friend Request from the Client to Recipient
func handleFriendRequestSent(userId string, message MessageWS) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	// Check the User Number of Friend Requests
	friend, err := services.GetUserById(message.RecipientID)
	if err != nil {
		return
	}
	numFriendRequests := strconv.Itoa(len(friend.FriendRequests))

	// Convert Friend Request into JSON
	friendRequest := MessageWS{
		Type:        "friend-requests",
		SenderID:    userId,
		RecipientID: message.RecipientID,
		Data:        numFriendRequests,
	}
	jsonFriendReq, err := json.Marshal(friendRequest)
	if err != nil {
		log.Println(err)
	}

	// Find a Recipient in the Connections
	client, ok := clients[message.RecipientID]
	if !ok {
		log.Println("Recipient not connected yet. ID:", message.RecipientID)
		return
	}

	// Send Message to the Recipient
	err = client.WriteMessage(websocket.TextMessage, jsonFriendReq)
	if err != nil {
		log.Println("WS Broadcast WriteMessage error:", err)
	}
}
