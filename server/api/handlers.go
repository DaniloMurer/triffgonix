package handlers

import (
	"net/http"
	"server/core/domain"
	engine2 "server/dart/engine"
	"server/dart/engine/x01"
	"server/database"

	"github.com/gin-gonic/gin"
)

var games []engine2.Game

func CreatePlayer(c *gin.Context) {
	player := domain.Player{Id: 0, PlayerName: "test"}
	database.CreatePlayer(player.ToPlayerEntity())
	c.JSON(http.StatusAccepted, gin.H{"text": "hello"})
}

func GetPlayers(c *gin.Context) {
	users := database.FindAllUsers()
	c.JSON(http.StatusFound, &users)
}

func CreateGame(c *gin.Context) {
	// TODO: implement game creation through post request
	game := engine2.Game{Name: "test", Engine: x01.New(301)}
	games = append(games, game)
}
