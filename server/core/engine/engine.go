package engine

import "server/core/domain"

type Player struct {
	Value    *domain.Player
	Previous *Player
	Next     *Player
}

// Players is a linked list of the players in a given game
type Players struct {
	Head          *Player
	CurrentPlayer *Player
	Tail          *Player
}

func (players *Players) Add(player *Player) {
	if players.Head == nil {
		players.Head = player
		players.Tail = player
	} else {
		players.Tail.Next = player
		player.Previous = players.Tail
		players.Tail = player
	}
}

func (players *Players) NextPlayer() *Player {
	nextPlayer := players.CurrentPlayer.Next
	if nextPlayer == nil {
		nextPlayer = players.Head
	}
	players.CurrentPlayer = nextPlayer
	return nextPlayer
}

// FIXME: maybe the engine should be part of the game struct for clarity sake
type Game struct {
	Name    string
	Players Players
	Throws  []domain.Throw
}

type Engine interface {
	// getPlayerThrows returns the throws made by a given player
	getPlayerThrows(player *domain.Player) *[]domain.Throw
	// nextPlayer returns the domain object of the next player and updates the linked list accordingly
	nextPlayer() *domain.Player
	// RegisterThrow registers a new player's throw and returns the player id of the next player
	RegisterThrow(throw *domain.Throw) uint
	// UndoThrow removes the last throw and returns the player id of the next player
	UndoThrow(throw *domain.Throw) uint
}
