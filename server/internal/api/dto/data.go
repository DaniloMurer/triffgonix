package dto

import (
	"github.com/DaniloMurer/triffgonix/server/internal/domain"
	"github.com/DaniloMurer/triffgonix/server/internal/models"
)

type Players struct {
	AllPlayers    []domain.Player `json:"allPlayers"`
	CurrentPlayer domain.Player   `json:"currentPlayer"`
}

type Game struct {
	Id            uint     `json:"id"`
	Name          string   `json:"name"`
	GameMode      string   `json:"gameMode"`
	StartingScore int16    `json:"startingScore"`
	Players       []Player `json:"players"`
}

func (game *Game) ToEntity() *models.Game {
	var players []models.Player
	for _, player := range game.Players {
		players = append(players, *player.ToEntity())
	}
	return &models.Game{
		Name:     game.Name,
		GameMode: game.GameMode,
		Players:  players,
	}
}

func (game *Game) FromEntity(gameEntity *models.Game) {
	var players []Player
	for _, player := range gameEntity.Players {
		players = append(players, Player{Id: player.Id, Name: player.PlayerName})
	}
	game.Id = gameEntity.Id
	game.Name = gameEntity.Name
	game.GameMode = gameEntity.GameMode
	game.Players = players
}

type Player struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (player Player) ToEntity() *models.Player {
	return &models.Player{
		Model:      models.Model{Id: player.Id},
		PlayerName: player.Name,
	}
}

func (player Player) ToDomain() *domain.Player {
	return &domain.Player{
		Id:            player.Id,
		PlayerName:    player.Name,
		Score:         0,
		AveragePoints: 0,
	}
}

func (player Player) FromEntity(playerEntity *models.Player) Player {
	return Player{
		Id:   playerEntity.Id,
		Name: playerEntity.PlayerName,
	}
}
