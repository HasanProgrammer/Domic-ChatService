package WebAPIRequest

import (
	"Domic.Domain/Commons/Contracts"
	"Domic.UseCase/ChatUseCase/DTOs"
	"Domic.UseCase/Commons/Contracts"
	"Domic.UseCase/UserUseCase/Commands"
	"Domic.WebAPI/DTOs"
	"Domic.WebAPI/Requests/Helpers"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type ChatRequestController struct {
	messageBroker    UseCaseCommonContract.IMessageBroker
	idGenerator      DomainCommonContract.IGlobalIdentityGenerator
	serializer       DomainCommonContract.ISerializer
	distributedCache UseCaseCommonContract.IInternalDistributedCache
	clients          map[string]*websocket.Conn
}

// SignInHandler concurrent runing
func (controller *ChatRequestController) SignInHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		fullName := r.PostFormValue("FullName")
		email := r.PostFormValue("EMail")

		signInCommandHandler, errs := UserUseCaseCommand.NewSignInCommandHandler(controller.distributedCache, fullName, email)

		responseObject := WebAPIDTO.JsonResponseDto{}

		if len(errs) > 0 {

			responseObject.Code = http.StatusBadRequest
			responseObject.Message = "ورود شما به پنل پشتیبانی با موفقیت انجام نگرفت!"
			responseObject.Body = errs

			WebAPIRequestHelper.WriteJsonResponse(controller.serializer, w, responseObject)

		} else {

			result := signInCommandHandler.Handle()

			if result {

				responseObject.Code = http.StatusOK
				responseObject.Message = "ورود شما به پنل پشتیبانی با موفقیت انجام شد"
				responseObject.Body = struct {
				}{}

				WebAPIRequestHelper.WriteJsonResponse(controller.serializer, w, responseObject)

			} else {

				responseObject.Code = http.StatusBadRequest
				responseObject.Message = "ورود شما به پنل پشتیبانی با موفقیت انجام نگرفت!"
				responseObject.Body = struct{}{}

				WebAPIRequestHelper.WriteJsonResponse(controller.serializer, w, responseObject)

			}

		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

// WsConnectionsHandler concurrent runing
func (controller *ChatRequestController) WsConnectionsHandler(w http.ResponseWriter, r *http.Request) {

	connectionId := controller.idGenerator.Generate()

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			//in here must be checked client url and then return false or true
			//this result is temporally!
			return true
		},
	}

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}

	defer ws.Close()

	controller.clients[connectionId] = ws

	fmt.Printf("User %s connected\n", connectionId)

	//long runing goroutin ( http goroutin )
	for {
		var chatDto ChatUseCaseDTO.ChatDto
		err := ws.ReadJSON(&chatDto)
		if err != nil {
			delete(controller.clients, connectionId)
			break
		}
		chatDto.ConnectionId = connectionId
		//produce message into broker
		controller.messageBroker.Publish(chatDto, "chat")
	}
}

func (controller *ChatRequestController) ConsumeChatMessagesHandler() {

	controller.messageBroker.Subscribe("chat", func(body []byte) error {

		message := ChatUseCaseDTO.ChatDto{}

		controller.serializer.Deserialize(string(body), &message)

		if message.To == "" {
			//send message to all clients
			for _, client := range controller.clients {
				client.WriteJSON(message)
			}
		} else {
			//send message to specific client
			if client, ok := controller.clients[message.To]; ok {
				client.WriteJSON(message)
			}
		}

		return nil

	})

}

func NewChatRequestController(broker UseCaseCommonContract.IMessageBroker, serializer DomainCommonContract.ISerializer,
	cache UseCaseCommonContract.IInternalDistributedCache, idGenerator DomainCommonContract.IGlobalIdentityGenerator,
) *ChatRequestController {
	return &ChatRequestController{
		messageBroker:    broker,
		serializer:       serializer,
		distributedCache: cache,
		idGenerator:      idGenerator,
		clients:          make(map[string]*websocket.Conn),
	}
}
