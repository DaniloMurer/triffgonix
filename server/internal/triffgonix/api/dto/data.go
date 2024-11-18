package dto

import (
	"server/internal/triffgonix/domain"
	"server/internal/triffgonix/models"
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

func (self *Game) ToEntity() *models.Game {
	var players []models.Player
	for _, player := range self.Players {
		players = append(players, *player.ToEntity())
	}
	return &models.Game{
		Name:    self.Name,
		Players: players,
	}
}

func (self *Game) FromEntity(game *models.Game) {
	var players []Player
	for _, player := range game.Players {
		players = append(players, Player{Id: player.Id, Name: player.PlayerName})
	}
	self.Id = game.Id
	self.Name = game.Name
	self.Players = players
}

type Player struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (self Player) ToEntity() *models.Player {
	return &models.Player{
		Model:      models.Model{Id: self.Id},
		PlayerName: self.Name,
	}
}

func (self Player) ToDomain() *domain.Player {
	return &domain.Player{
		Id:            self.Id,
		PlayerName:    self.Name,
		Score:         0,
		AveragePoints: 0,
	}
}
