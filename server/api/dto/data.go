package dto

import (
	"server/core/domain"
	"server/database"
)

type Players struct {
	AllPlayers    []domain.Player `json:"allPlayers"`
	CurrentPlayer domain.Player   `json:"currentPlayer"`
}

type Game struct {
	Name          string   `json:"name"`
	GameMode      string   `json:"gameMode"`
	StartingScore int16    `json:"startingScore"`
	Players       []Player `json:"players"`
}

func (self Game) ToEntity() *database.Game {
	var players []database.Player
	for _, player := range self.Players {
		players = append(players, player.ToEntity())
	}
	return &database.Game{
		Name:    self.Name,
		Players: players,
	}
}

type Player struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (self Player) ToEntity() database.Player {
	return database.Player{
		Model:      database.Model{Id: self.Id},
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
