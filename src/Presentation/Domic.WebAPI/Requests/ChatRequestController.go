package WebAPIRequest

import (
	"Domic.Domain/Commons/Contracts"
	"Domic.UseCase/Commons/Contracts"
	"Domic.UseCase/UserUseCase/Commands"
	"Domic.WebAPI/DTOs"
	"Domic.WebAPI/Requests/Helpers"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var broadcast = make(chan ChatDto)

type ChatDto struct {
	ConnectionId string `json:"connectionId"`
	Content      string `json:"content"`
	To           string `json:"to"`
}

type ChatRequestController struct {
	serializer       DomainCommonContract.ISerializer
	distributedCache UseCaseCommonContract.IInternalDistributedCache
	clients          map[string]*websocket.Conn
}

func (controller *ChatRequestController) SignInAction(w http.ResponseWriter, r *http.Request) {

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
				responseObject.Body = struct {
				}{}

				WebAPIRequestHelper.WriteJsonResponse(controller.serializer, w, responseObject)

			}

		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (controller *ChatRequestController) WsConnectionsAction(w http.ResponseWriter, r *http.Request) {

	connectionId := r.URL.Query().Get("ConnectionId")

	fmt.Println(connectionId, nil)

	if connectionId == "" {
		http.Error(w, "Missing [ConnectionId]!", http.StatusBadRequest)
		return
	}

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

	for {
		var chatDto ChatDto
		err := ws.ReadJSON(&chatDto)
		if err != nil {
			delete(controller.clients, connectionId)
			break
		}
		chatDto.ConnectionId = connectionId
		broadcast <- chatDto
	}
}

func (controller *ChatRequestController) ConsumeMessagesAction() {
	for {
		msg := <-broadcast
		if msg.To == "" {
			// Broadcast به همه
			for _, client := range controller.clients {
				client.WriteJSON(msg)
			}
		} else {

			// ارسال به گیرنده خاص
			if client, ok := controller.clients[msg.To]; ok {
				client.WriteJSON(msg)
			}
		}
	}
}

func NewChatRequestController() *ChatRequestController {
	return &ChatRequestController{
		clients: make(map[string]*websocket.Conn),
	}
}
