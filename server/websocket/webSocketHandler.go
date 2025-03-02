package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type MessageWS struct {
	Type       string `json:"type"`
	SenderID   int    `json:"senderId,omitempty"`
	ReceiverID int    `json:"receiverId,omitempty"`
	Data       string `json:"data"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var clients = make(map[int]*websocket.Conn)
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
	id, _ := strconv.Atoi(senderId)

	// Add connected Sender ID as a Client
	addClient(id, ws)

	// Read Incoming Messages
	go reader(id, ws)
}

func addClient(clientId int, conn *websocket.Conn) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	clients[clientId] = conn
	log.Println(clientId, "connected")
}

func reader(userId int, conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("ReadMessage error:", err)
			clients[userId].Close()
			delete(clients, userId)
			log.Println("Client closed and removed: ", userId)
			return
		}

		log.Println("Reader msg: ", string(p))
		jsonData := string(p)

		var message MessageWS
		err = json.Unmarshal([]byte(jsonData), &message)
		if err != nil {
			log.Println("WS Reader error parsing JSON: ", err)
		}

		if message.Type == "chat-connection" {
			handleChatConnection(userId, message)
		} else if message.Type == "chat-message" {
			handleChatMessage(userId, message)
		}
	}
}

func handleChatConnection(userId int, message MessageWS) {
	// TEMP:
	// Broadcast the message for all clients besides the active one
	msg := MessageWS{Type: message.Type, SenderID: userId, Data: "joined to the chat"}

	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	for clientId, client := range clients {
		// TEMP:
		// Match only the clients besides the sender
		if clientId == userId {
			continue
		}

		log.Println("sending to:", clientId)

		err := client.WriteMessage(websocket.TextMessage, jsonMsg)
		if err != nil {
			log.Println("WS Broadcast WriteMessage error:", err)
			// client.Close()
			// delete(clients, userId)
		}

	}
}

func handleChatMessage(userId int, message MessageWS) {
	// TEMP:
	// Broadcast the message for all clients besides the active one
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	jsonMsg, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	for clientId, client := range clients {
		// TEMP:
		// Match only the clients besides the sender
		if clientId == userId {
			continue
		}

		log.Println("sending to:", clientId)

		err := client.WriteMessage(websocket.TextMessage, jsonMsg)
		if err != nil {
			log.Println("WS Broadcast WriteMessage error:", err)
			// client.Close()
			// delete(clients, userId)
		}

	}
}
