package Models

import (
	"gorm.io/gorm"
	"time"
)

type ChatModel struct {
	gorm.Model

	Id           string     `gorm:"primaryKey" gorm:"column:Id"`
	UserId       string     `gorm:"column:UserId"`
	ConnectionId string     `gorm:"column:ConnectionId"`
	Content      string     `gorm:"column:Content"`
	To           string     `gorm:"column:To"` //other side user = other [ ConnectionId ]
	CreatedBy    string     `gorm:"column:CreatedBy"`
	CreatedRole  string     `gorm:"column:CreatedRole"`
	CreatedAt    time.Time  `gorm:"column:CreatedAt"`
	UpdatedBy    *string    `gorm:"column:UpdatedBy"`
	UpdatedRole  *string    `gorm:"column:UpdatedRole"`
	UpdatedAt    *time.Time `gorm:"column:UpdatedAt"`
	Version      string     `gorm:"column:Version"`
	IsDeleted    bool       `gorm:"column:IsDeleted"`

	User UserModel
}
