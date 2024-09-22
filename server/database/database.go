package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func openDatabaseConnection() {
	db, err := gorm.Open(sqlite.Open("triffnix.db"), &gorm.Config{})
	if err != nil {
		panic("error while opening database connection")
	}
	database = db
}

func AutoMigrate() {
	openDatabaseConnection()
	err := database.AutoMigrate(&Player{})
	if err != nil {
		panic("cannot migrate schema to database")
	}
}

func FindAllUsers() []Player {
	var users []Player
	database.Find(&users)
	return users
}

func CreatePlayer(player *Player) {
	database.Save(player)
}

func CreateDummyUser() {
	user := Player{PlayerName: "testico"}
	database.Save(&user)
}
