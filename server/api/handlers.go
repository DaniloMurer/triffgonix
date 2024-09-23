package handlers

import (
	"net/http"
	"server/core/domain"
	"server/core/engine"
	"server/database"

	"github.com/gin-gonic/gin"
)

var games []engine.Game

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
	game := engine.Game{Name: "test", Engine: &engine.X01Engine{}}
	game.Engine.NextPlayer(game.Players)
	games = append(games, game)
}
