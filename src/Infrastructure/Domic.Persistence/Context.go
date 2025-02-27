package DomicPersistence

import (
	"Domic.Persistence/Models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DbContext struct {
	postgresDb *gorm.DB
}

func NewDbContext(connection string) *DbContext {
	db, err := gorm.Open(sqlserver.Open(connection), &gorm.Config{})

	if err != nil {
	}

	db.AutoMigrate(&Models.ChatModel{})
	db.AutoMigrate(&Models.UserModel{})
	db.AutoMigrate(&Models.EventModel{})

	return &DbContext{postgresDb: db}
}
