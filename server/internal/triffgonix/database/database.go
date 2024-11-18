package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"server/internal/triffgonix/models"
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
	err := database.AutoMigrate(&models.Player{}, &models.Game{})
	if err != nil {
		panic("cannot migrate schema to database")
	}
}

func FindAllUsers() []models.Player {
	var users []models.Player
	database.Find(&users)
	return users
}

func CreatePlayer(player *models.Player) (error, *models.Player) {
	var err error
	result := database.Save(player)
	if result.Error != nil {
		err = result.Error
	}
	return err, player
}

func FindAllGames() []models.Game {
	var games []models.Game
	database.Find(&games)
	return games
}

func CreateGame(game *models.Game) (error, *models.Game) {
	result := database.Save(game)
	var err error
	if result.Error != nil {
		err = result.Error
	}
	return err, game
}

func CreateDummyUser() {
	user := models.Player{PlayerName: "testico"}
	database.Save(&user)
}
