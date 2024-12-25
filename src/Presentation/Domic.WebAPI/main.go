package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {

		//in here must be checked client url and then return false or true
		//this result is temporally!
		return true

	},
}

var clients = make(map[string]*websocket.Conn)
var broadcast = make(chan Message)

type Message struct {
	ConnectionId string `json:"connectionId"`
	Content      string `json:"content"`
	To           string `json:"to"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	connectionId := r.URL.Query().Get("connectionId")

	if connectionId == "" {
		http.Error(w, "Missing [ConnectionId]!", http.StatusBadRequest)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}

	defer ws.Close()

	clients[connectionId] = ws
	fmt.Printf("User %s connected\n", connectionId)

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, connectionId)
			break
		}
		msg.ConnectionId = connectionId
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		if msg.To == "" {
			// Broadcast به همه
			for _, client := range clients {
				client.WriteJSON(msg)
			}
		} else {

			// ارسال به گیرنده خاص
			if client, ok := clients[msg.To]; ok {
				client.WriteJSON(msg)
			}
		}
	}
}

func main() {

	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
