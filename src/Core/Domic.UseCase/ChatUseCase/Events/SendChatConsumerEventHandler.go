package ChatUseCaseEvent

import (
	"Domic.Domain/Chat/Contracts"
	"Domic.Domain/Chat/Entities"
	"Domic.Domain/Commons/Contracts"
	"Domic.Domain/Commons/Entities"
	"Domic.UseCase/ChatUseCase/DTOs"
)

type SendChatConsumerEventHandler struct {
	idGenerator           DomainCommonContract.IGlobalIdentityGenerator
	serializer            DomainCommonContract.ISerializer
	chatCommandRepository DomainChatContract.IChatCommandRepository
}

func (consumer *SendChatConsumerEventHandler) Handle(event *DomainCommonEntity.Event) error {

	chatMessage := &ChatUseCaseDTO.ChatDto{}

	err := consumer.serializer.Deserialize(event.GetPayload(), chatMessage)

	if err != nil {
		return err
	}

	newChatEntity, entityError := DomainChatEntity.NewChat(
		consumer.idGenerator,
		chatMessage.ConnectionId,
		chatMessage.Content,
		chatMessage.To,
		event.GetCreatedBy(),
		event.GetCreatedRole(),
	)

	if entityError != nil {
		return entityError
	}

	consumer.chatCommandRepository.Add(newChatEntity)

	return nil
}

func NewSendChatConsumerEventHandler(chatCommandRepository DomainChatContract.IChatCommandRepository,
	serializer DomainCommonContract.ISerializer,
) *SendChatConsumerEventHandler {
	return &SendChatConsumerEventHandler{
		serializer:            serializer,
		chatCommandRepository: chatCommandRepository,
	}
}
