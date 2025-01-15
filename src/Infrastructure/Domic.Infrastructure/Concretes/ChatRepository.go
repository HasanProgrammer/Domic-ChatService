package InfrastructureConcrete

import (
	"Domic.Domain/Chat/Entities"
	"Domic.Domain/Commons/DTOs"
	"Domic.Persistence/Models"
	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func (ChatRepository *ChatRepository) Add(entity *DomainChatEntity.Chat) DomainCommonDTO.Result[bool] {

	model := Models.ChatModel{}

	queryResult := ChatRepository.db.Create(&model)

	return DomainCommonDTO.Result[bool]{
		Error:  queryResult.Error,
		Result: queryResult.Error == nil,
	}

}

func (ChatRepository *ChatRepository) AddRange(entities *[]DomainChatEntity.Chat) DomainCommonDTO.Result[bool] {

	queryResult := ChatRepository.db.CreateInBatches(entities, len(*entities))

	return DomainCommonDTO.Result[bool]{
		Error:  queryResult.Error,
		Result: queryResult.Error != nil,
	}

}

func (ChatRepository *ChatRepository) Change(entity *DomainChatEntity.Chat) DomainCommonDTO.Result[bool] {

	queryResult := ChatRepository.db.Save(entity)

	return DomainCommonDTO.Result[bool]{
		Error:  queryResult.Error,
		Result: queryResult.Error != nil,
	}
}

func (ChatRepository *ChatRepository) Remove(entity *DomainChatEntity.Chat) DomainCommonDTO.Result[bool] {

	queryResult := ChatRepository.db.Delete(entity)

	return DomainCommonDTO.Result[bool]{
		Error:  queryResult.Error,
		Result: queryResult.Error != nil,
	}
}

func (ChatRepository *ChatRepository) FindById(id string) DomainCommonDTO.Result[*DomainChatEntity.Chat] {

	var Chat *DomainChatEntity.Chat

	queryResult := ChatRepository.db.First(Chat, "id = ?", id)

	return DomainCommonDTO.Result[*DomainChatEntity.Chat]{
		Error:  queryResult.Error,
		Result: Chat,
	}

}

func (ChatRepository *ChatRepository) FindAllPaginated(paginationRequest *DomainCommonDTO.PaginationRequest) DomainCommonDTO.Result[DomainCommonDTO.PaginationResponse[DomainChatEntity.Chat]] {

	var Chats *[]DomainChatEntity.Chat

	take := int(paginationRequest.PageSize)
	skip := int((paginationRequest.PageIndex - 1) * paginationRequest.PageSize)

	queryResult := ChatRepository.db.Limit(take).Offset(skip).Find(&Chats)

	return DomainCommonDTO.Result[DomainCommonDTO.PaginationResponse[DomainChatEntity.Chat]]{
		Error: queryResult.Error,
		Result: DomainCommonDTO.PaginationResponse[DomainChatEntity.Chat]{
			PageSize:  paginationRequest.PageSize,
			PageIndex: paginationRequest.PageIndex,
			Items:     Chats,
		},
	}
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{db: db}
}
