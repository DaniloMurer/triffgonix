package dto

import (
	"server/core/domain"
)

type Players struct {
	AllPlayers    []domain.Player `json:"allPlayers"`
	CurrentPlayer domain.Player   `json:"currentPlayer"`
}
