package main

import (
	"Domic.WebAPI/Requests"
	"fmt"
	"net/http"
)

func main() {

	chatRequestController := WebAPIRequest.NewChatRequestController()

	http.HandleFunc("/chat.css", WebAPIRequest.HandleStyle)
	http.HandleFunc("/chat.js", WebAPIRequest.HandleScript)
	http.HandleFunc("/chat-ui", WebAPIRequest.HandlePublicChatPage)
	http.HandleFunc("/signin", chatRequestController.SignInAction)
	http.HandleFunc("/chat", chatRequestController.WsConnectionsAction)

	go chatRequestController.ConsumeMessagesAction()

	fmt.Println("WebSocket server started on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
