package websocket

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Message struct {
	Type string `json:"type"`
	Data int    `json:"data"`
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
			return
		}

		log.Println("Reader userId: ", userId)
		log.Println("Reader msg: ", string(p))

		// if string(p) == "userMessages" {
		// 	handleSentMessage()
		// } else if string(p) == "messageSent" {
		// 	handleSentMessage()
		// }
	}
}

// func handleSentMessage() {
// 	msg := Message{Type: "incomingMessage", Data: 1}
// 	broadcastMessage(msg)
// }

// func broadcastMessage(message Message) {
// 	clientsMutex.Lock()
// 	defer clientsMutex.Unlock()

// 	msg, err := json.Marshal(message)
// 	if err != nil {
// 		log.Println("WS Broadcast error encoding message:", err)
// 		return
// 	}

// 	for userId, client := range clients {
// 		if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
// 			log.Println("WS Broadcast WriteMessage error:", err)
// 			client.Close()
// 			delete(clients, userId)
// 		}

// 	}
// }
