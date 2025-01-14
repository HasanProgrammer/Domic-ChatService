package ChatEntity

import (
	"Domic.Domain/Commons/Contracts"
	"errors"
	"time"
)

type Chat struct {
	id           string
	userId       string
	connectionId string
	content      string
	to           string //other side user = other [ ConnectionId ]
	createdBy    string
	createdRole  string
	createdAt    time.Time
	version      string
	isDeleted    bool
}

func NewChat(idGenerator DomainCommonContract.IGlobalIdentityGenerator, serializer DomainCommonContract.ISerializer,
	identity DomainCommonContract.IIdentityUser, connectionId string, content string, to string,
) (*Chat, error) {

	if len(content) > 1000 {
		return nil, errors.New("متن پیام ارسالی نباید بیشتر از 1000 عبارت داشته باشد!")
	}

	uniqueId := idGenerator.Generate()
	nowTime := time.Now()
	userId := identity.GetUserIdentity()
	userRoles, err := serializer.Serialize(identity.GetUserRoles())

	if err != nil {
		return nil, errors.New("متن پیام ارسالی نباید بیشتر از 1000 عبارت داشته باشد!")
	}

	newChat := Chat{
		id:           uniqueId,
		userId:       userId,
		connectionId: connectionId,
		content:      content,
		to:           to,
		createdBy:    userId,
		createdRole:  userRoles,
		createdAt:    nowTime,
		version:      idGenerator.Generate(),
	}

	return &newChat, nil

}
