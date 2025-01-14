package DomainUserEntity

import (
	"Domic.Domain/Chat/Entities"
	"time"
)

type User struct {
	Id          string
	FullName    string
	Email       string
	CreatedRole string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedRole *string
	UpdatedBy   *string
	UpdatedAt   *time.Time
	Version     string
	IsDeleted   bool
	IsActive    bool
	Chats       *[]DomainChatEntity.Chat
}
