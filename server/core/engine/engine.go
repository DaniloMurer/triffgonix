package engine

import (
	"fmt"
	"server/core/domain"
)

type Engine interface {
	// GetPlayerThrows returns the throws made by a given player
	GetPlayerThrows(player *Player) *[]domain.Throw
	// NextPlayer returns the domain object of the next player and updates the linked list accordingly
	NextPlayer(players *Players) *domain.Player
	// RegisterThrow registers a new player's throw
	RegisterThrow(throw *domain.Throw, players *Players)
	// UndoThrow removes the last throw
	UndoThrow(throw *domain.Throw, players *Players)
	// CalculatePlayerScore returns the player score across all turns
	CalculatePlayerScore(player *Player, startingScore uint16)
}

type Player struct {
	Value    *domain.Player
	Previous *Player
	Next     *Player
	Turns    []Turn
}

func (player *Player) CalculateScore(startingScore uint16) {
	var totalSum uint16
	for _, turn := range player.Turns {
		totalSum += uint16(turn.Sum())
	}
	player.Value.Score = startingScore - totalSum
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
	First  *domain.Throw
	Second *domain.Throw
	Third  *domain.Throw
}

func (turn Turn) Sum() uint8 {
	var first uint8
	var second uint8
	var third uint8
	fmt.Printf("trace: turn - %+v\n", turn)
	if turn.First != nil {
		first = turn.First.Points * turn.First.Multiplicator
	}
	if turn.Second != nil {
		second = turn.Second.Points * turn.Second.Multiplicator
	}
	if turn.Third != nil {
		third = turn.Third.Points * turn.Third.Multiplicator
	}

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

func (turn Turn) HasSpace() bool {
	if turn.First != nil && turn.Second != nil && turn.Third != nil {
		return false
	} else {
		return true
	}
}

type Game struct {
	Name          string
	StartingScore uint16
	Players       *Players
	Throws        *[]domain.Throw
	Engine        Engine
}
