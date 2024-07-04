package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-discord-clone/models"
)

var DBConn *gorm.DB

func Connect() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("./database/discord.db"))
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	DBConn.AutoMigrate(
		&models.Community{},
		&models.User{},
	)
}
