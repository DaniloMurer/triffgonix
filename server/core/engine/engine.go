package engine

import "server/core/domain"

type Player struct {
	Value    domain.Player
	Previous domain.Player
	Next     domain.Player
}

// Players is a linked list of the players in a given game
type Players struct {
	Head          Player
	CurrentPlayer Player
}

// TODO: implement linked list

type Game struct {
	Name    string
	Players Players
	Throws  []domain.Throw
}

type Engine interface {
	// getPlayerThrows returns the throws made by a given player
	getPlayerThrows(player *domain.Player) *[]domain.Throw
	// getNextPlayer returns the domain object of the next player and updates the linked list accordingly
	getNextPlayer(player *Player) *domain.Player
	// RegisterThrow registers a new player's throw and returns the player id of the next player
	RegisterThrow(throw *domain.Throw) uint
	// UndoThrow removes the last throw and returns the player id of the next player
	UndoThrow(throw *domain.Throw) uint
}
