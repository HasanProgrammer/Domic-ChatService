package UserContract

import (
	"Domic.Domain/Commons/Contracts"
	"Domic.Domain/User/Entities"
)

type IUserCommandRepository interface {
	DomainCommonContract.IRepository[string, DomainUserEntity.User]
}
