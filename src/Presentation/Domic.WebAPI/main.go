package main

import (
	"Domic.Infrastructure/Concretes"
	"Domic.WebAPI/Requests"
	"fmt"
	"net/http"
)

func main() {

	configuration := InfrastructureConcrete.NewConfiguration()
	serializer := InfrastructureConcrete.NewSerializer()

	rabbitConnectionString, err := configuration.GetConnectionString("I-RabbitMQ")
	redisConnectionString, err := configuration.GetConnectionString("I-Redis")

	chatRequestController := WebAPIRequest.NewChatRequestController(
		InfrastructureConcrete.NewMessageBroker(rabbitConnectionString),
		serializer,
		InfrastructureConcrete.NewDistributedCache(serializer, redisConnectionString, ""),
		InfrastructureConcrete.NewGlobalIdentityGenerator(),
	)

	//sync requests
	http.HandleFunc("/chat.css", WebAPIRequest.HandleStyle)
	http.HandleFunc("/chat.js", WebAPIRequest.HandleScript)
	http.HandleFunc("/chat-ui", WebAPIRequest.HandlePublicChatPage)
	http.HandleFunc("/signin", chatRequestController.SignInHandler)
	http.HandleFunc("/chat", chatRequestController.WsConnectionsHandler)

	//async requests
	chatRequestController.ConsumeChatMessagesHandler()

	//start server listener

	fmt.Println("WebSocket server started on :8080")

	serverError := http.ListenAndServe(":8080", nil)

	if serverError != nil {
		fmt.Println("Error starting server:", err)
	}

}
