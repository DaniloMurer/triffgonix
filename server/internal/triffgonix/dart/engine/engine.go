package engine

import (
	"server/internal/triffgonix/api/dto"
	"server/internal/triffgonix/domain"
	"server/pkg/logging"
)

var logger logging.Logger = logging.NewLogger()

type Engine interface {
	// FIXME: i think for the future it would make sense to have a way to tell if certain points can be made.
	// an example would be shangai, where after hitting three times a certain point, you cannot score on it no more.

	// GetPlayerThrows returns the throws made by a given player
	GetPlayerThrows(player *Player) *[]domain.Throw
	// RegisterThrow registers a new player's throw
	RegisterThrow(throw *domain.Throw, players *Players)
	// UndoLastThrow removes the last throw
	UndoLastThrow(players *Players)
	// CalculatePlayerScore returns the player score across all turns
	CalculatePlayerScore(player *Player)
	// HasAnyPlayerWon returns then winning player if one exists
	HasAnyPlayerWon(players *Players) *Player
}

type Player struct {
	Value    *domain.Player
	Previous *Player
	Next     *Player
	Turns    []Turn
}

// GetAveragePoints gets average points scored across all turns from player
func (player *Player) GetAveragePoints() int16 {
	var averagePoints int16
	var throwCount int16
	for _, turn := range player.Turns {
		averagePoints += turn.Sum()
		throwCount += turn.ThrowCount()
	}
	if throwCount == 0 {
		return 0
	}
	return averagePoints / throwCount
}

// Players is a linked list of the players in a given game
type Players struct {
	Head          *Player
	CurrentPlayer *Player
	Tail          *Player
}

// Add adds new player to the linked list
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

// SwitchToNextPlayer switches to the next player from the CurrentPlayer perspective
func (players *Players) SwitchToNextPlayer() *Player {
	nextPlayer := players.CurrentPlayer.Next
	if nextPlayer == nil {
		nextPlayer = players.Head
	}
	players.CurrentPlayer = nextPlayer
	return nextPlayer
}

// SwitchToPreviousPlayer switches to the previous player from the CurrentPlayer perspective
func (players *Players) SwitchToPreviousPlayer() *Player {
	previousPlayer := players.CurrentPlayer.Previous
	if previousPlayer == nil {
		previousPlayer = players.Tail
	}
	players.CurrentPlayer = previousPlayer
	return previousPlayer
}

// GetPreviousPlayer returns the previous player from the CurrentPlayer perspective
func (players *Players) GetPreviousPlayer() *Player {
	previousPlayer := players.CurrentPlayer.Previous
	if previousPlayer == nil {
		previousPlayer = players.Tail
	}
	return previousPlayer
}

func (players *Players) ToDto() dto.Players {
	dtoPlayers := dto.Players{}
	var domainPlayers []domain.Player
	currentNode := players.Head
	for currentNode != nil {
		domainPlayer := *currentNode.Value
		domainPlayer.AveragePoints = currentNode.GetAveragePoints()
		domainPlayers = append(domainPlayers, domainPlayer)
		currentNode = currentNode.Next
	}
	dtoPlayers.AllPlayers = domainPlayers
	dtoPlayers.CurrentPlayer = *players.CurrentPlayer.Value
	return dtoPlayers
}

type Turn struct {
	First  *domain.Throw
	Second *domain.Throw
	Third  *domain.Throw
}

// Sum returns the sum of all throws in the turn
func (turn *Turn) Sum() int16 {
	var first int16
	var second int16
	var third int16
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

// HasSpace checks if the turn has a free slot for a throw
func (turn *Turn) HasSpace() bool {
	if turn.First != nil && turn.Second != nil && turn.Third != nil {
		return false
	} else {
		return true
	}
}

// FillTurn fills the turn with 0 points throws
func (turn *Turn) FillTurn(throw *domain.Throw) {
	if turn.First == nil {
		turn.First = &domain.Throw{Points: 0, Multiplicator: 1, PlayerId: throw.PlayerId}
	}
	if turn.Second == nil {
		turn.Second = &domain.Throw{Points: 0, Multiplicator: 1, PlayerId: throw.PlayerId}
	}
	if turn.Third == nil {
		turn.Third = &domain.Throw{Points: 0, Multiplicator: 1, PlayerId: throw.PlayerId}
	}
}

// ThrowCount returns the count of throws in the turn
func (turn *Turn) ThrowCount() int16 {
	var throwCount int16 = 0
	if turn.First != nil {
		throwCount += 1
	}
	if turn.Second != nil {
		throwCount += 1
	}
	if turn.Third != nil {
		throwCount += 1
	}
	return throwCount
}

type Game struct {
	Name    string
	Players *Players
	Engine  Engine
}
