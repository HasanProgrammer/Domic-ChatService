package Models

import "time"

type UserModel struct {
	Id          string    `gorm:"primaryKey" gorm:"column:Id"`
	FullName    string    `gorm:"column:FullName"`
	Email       string    `gorm:"column:Email"`
	CreatedBy   string    `gorm:"column:CreatedBy"`
	CreatedRole string    `gorm:"column:CreatedRole"`
	CreatedAt   time.Time `gorm:"column:CreatedAt"`
	version     string    `gorm:"column:version"`
	IsDeleted   bool      `gorm:"column:IsDeleted"`
	IsActive    bool      `gorm:"column:IsActive"`
}
