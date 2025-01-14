package Models

import "time"

type ChatModel struct {
	Id           string    `gorm:"primaryKey" gorm:"column:Id"`
	ConnectionId string    `gorm:"column:ConnectionId"`
	Content      string    `gorm:"column:Content"`
	To           string    `gorm:"column:To"` //other side user = other [ ConnectionId ]
	CreatedBy    string    `gorm:"column:CreatedBy"`
	CreatedRole  string    `gorm:"column:CreatedRole"`
	CreatedAt    time.Time `gorm:"column:CreatedAt"`
	Version      string    `gorm:"column:Version"`
	IsDeleted    bool      `gorm:"column:IsDeleted"`
}
