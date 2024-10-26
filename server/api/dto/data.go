package dto

import (
	"server/core/domain"
)

type Players struct {
	AllPlayers    []domain.Player `json:"allPlayers"`
	CurrentPlayer domain.Player   `json:"currentPlayer"`
}

type Game struct {
	Name          string   `json:"name"`
	GameMode      string   `json:"gameMode"`
	StartingScore int16    `json:"StartingScore"`
	Players       []Player `json:"players"`
}

type Player struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
