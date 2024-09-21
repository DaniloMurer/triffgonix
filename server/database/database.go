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
	err := database.AutoMigrate(&User{})
	if err != nil {
		panic("cannot migrate schema to database")
	}
}

func FindAllUsers() []User {
	var users []User
	database.Find(&users)
	return users
}

func CreateDummyUser() {
	user := User{UserName: "testico"}
	database.Save(&user)
}
