package DomainChatEntity

import (
	"time"
)

type StringifyChat struct {
	Id           string    `json:"Id"`
	UserId       string    `json:"UserId"`
	ConnectionId string    `json:"ConnectionId"`
	Content      string    `json:"Content"`
	To           string    `json:"To"`
	CreatedBy    string    `json:"CreatedBy"`
	CreatedRole  string    `json:"CreatedRole"`
	CreatedAt    time.Time `json:"CreatedAt"`
	Version      string    `json:"Version"`
	IsDeleted    bool      `json:"IsDeleted"`
}
