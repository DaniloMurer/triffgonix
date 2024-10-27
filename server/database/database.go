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
	err := database.AutoMigrate(&Player{}, &Game{})
	if err != nil {
		panic("cannot migrate schema to database")
	}
}

func FindAllUsers() []Player {
	var users []Player
	database.Find(&users)
	return users
}

func CreatePlayer(player *Player) (error, *Player) {
	var err error
	result := database.Save(player)
	if result.Error != nil {
		err = result.Error
	}
	return err, player
}

func FindAllGames() []Game {
	var games []Game
	database.Find(&games)
	return games
}

func CreateGame(game *Game) (error, *Game) {
	result := database.Save(game)
	var err error
	if result.Error != nil {
		err = result.Error
	}
	return err, game
}

func CreateDummyUser() {
	user := Player{PlayerName: "testico"}
	database.Save(&user)
}
