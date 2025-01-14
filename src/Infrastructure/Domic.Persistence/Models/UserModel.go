package Models

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	gorm.Model

	Id          string     `gorm:"primaryKey" gorm:"column:Id"`
	FullName    string     `gorm:"column:FullName"`
	Email       string     `gorm:"column:Email"`
	CreatedBy   string     `gorm:"column:CreatedBy"`
	CreatedRole string     `gorm:"column:CreatedRole"`
	CreatedAt   time.Time  `gorm:"column:CreatedAt"`
	UpdatedBy   *string    `gorm:"column:UpdatedBy"`
	UpdatedRole *string    `gorm:"column:UpdatedRole"`
	UpdatedAt   *time.Time `gorm:"column:UpdatedAt"`
	Version     string     `gorm:"column:version"`
	IsDeleted   bool       `gorm:"column:IsDeleted"`
	IsActive    bool       `gorm:"column:IsActive"`

	Chats []ChatModel `gorm:"foreignKey:UserId"`
}
