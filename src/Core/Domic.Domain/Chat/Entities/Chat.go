package DomainChatEntity

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

func NewChat(idGenerator DomainCommonContract.IGlobalIdentityGenerator, connectionId string, content string, to string,
	createdBy string, createdRoles string,
) (*Chat, error) {

	if len(content) > 1000 {
		return nil, errors.New("متن پیام ارسالی نباید بیشتر از 1000 عبارت داشته باشد!")
	}

	uniqueId := idGenerator.Generate()
	nowTime := time.Now()

	newChat := Chat{
		id:           uniqueId,
		userId:       createdBy,
		connectionId: connectionId,
		content:      content,
		to:           to,
		createdBy:    createdBy,
		createdRole:  createdRoles,
		createdAt:    nowTime,
		version:      idGenerator.Generate(),
	}

	return &newChat, nil

}
