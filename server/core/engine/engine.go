package engine

import (
	"server/core/domain"
)

type Engine interface {
	// GetPlayerThrows returns the throws made by a given player
	GetPlayerThrows(player *domain.Player, turns []Turn) *[]domain.Throw
	// NextPlayer returns the domain object of the next player and updates the linked list accordingly
	NextPlayer(players *Players) *domain.Player
	// RegisterThrow registers a new player's throw
	RegisterThrow(throw *domain.Throw, turns *[]Turn, players *Players)
	// UndoThrow removes the last throw
	UndoThrow(throw *domain.Throw, turns *[]Turn, players *Players)
}

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
		player.Previous = nil
		players.Head = player
		players.Tail = player
		// at the beginning, the head is always the current player
		players.CurrentPlayer = player
	} else {
		player.Next = nil
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

func (players *Players) PreviousPlayer() *Player {
	previousPlayer := players.CurrentPlayer.Previous
	if previousPlayer == nil {
		previousPlayer = players.Tail
	}
	players.CurrentPlayer = previousPlayer
	return previousPlayer
}

type Turn struct {
	First    *domain.Throw
	Second   *domain.Throw
	Third    *domain.Throw
	PlayerId uint
}

func (turn Turn) Sum() uint8 {
	first := turn.First.Points * turn.First.Multiplicator
	second := turn.Second.Points * turn.Second.Multiplicator
	third := turn.Third.Points * turn.Third.Multiplicator

	return first + second + third
}

// Append appends throw to turn. Returns true to signal a player switch is needed
func (turn *Turn) Append(throw *domain.Throw) bool {
	if turn.First == nil {
		turn.First = throw
		return false
	} else if turn.Second == nil {
		turn.Second = throw
		return false
	} else if turn.Third == nil {
		turn.Third = throw
		return true
	}
	return true
}

type Game struct {
	Name    string
	Players *Players
	Turns   []Turn
	Throws  *[]domain.Throw
	Engine  Engine
}
