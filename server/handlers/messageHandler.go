package handlers

import (
	"cosmic-gate-chat/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

// Get Messages from DB for all SenderID-RecipientID combinations
func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	senderId := r.URL.Query().Get("senderId")
	recipientId := r.URL.Query().Get("recipientId")

	messages, err := repositories.GetMessages(senderId, recipientId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting messages: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
