package DomainChatContract

import (
	"Domic.Domain/Chat/Entities"
	"Domic.Domain/Commons/Contracts"
)

type IChatCommandRepository interface {
	DomainCommonContract.IRepository[string, DomainChatEntity.Chat]
}
