package x01

import (
	"server/core/domain"
	"server/dart/engine"
	"server/logging"
)

var logger = logging.NewLogger()

type X01Engine struct {
	StartingScore int16
	Points        []int16
}

func (x01Engine *X01Engine) GetPlayerThrows(player *engine.Player) *[]domain.Throw {
	var playerThrows []domain.Throw
	for _, turn := range player.Turns {
		if turn.First != nil {
			playerThrows = append(playerThrows, *turn.First)
		}
		if turn.Second != nil {
			playerThrows = append(playerThrows, *turn.Second)
		}
		if turn.Third != nil {
			playerThrows = append(playerThrows, *turn.Third)
		}
	}
	return &playerThrows
}

func (x01Engine *X01Engine) RegisterThrow(throw *domain.Throw, players *engine.Players) {
	latestTurnIndex := len(players.CurrentPlayer.Turns) - 1
	// if player has no turns, then one should be created first
	if latestTurnIndex < 0 {
		newTurn := &engine.Turn{}
		hasOverthrown := x01Engine.checkForOverThrow(throw, players.CurrentPlayer)
		if hasOverthrown {
			newTurn.FillTurn(throw)
			players.CurrentPlayer.Turns = append(players.CurrentPlayer.Turns, *newTurn)
			x01Engine.CalculatePlayerScore(players.CurrentPlayer)
			players.SwitchToNextPlayer()
			return
		} else {
			newTurn.Append(throw)
			players.CurrentPlayer.Turns = append(players.CurrentPlayer.Turns, *newTurn)
			x01Engine.CalculatePlayerScore(players.CurrentPlayer)
			return
		}
	}
	latestTurn := &players.CurrentPlayer.Turns[latestTurnIndex]
	if !latestTurn.HasSpace() {
		newTurn := &engine.Turn{}
		hasOverthrown := x01Engine.checkForOverThrow(throw, players.CurrentPlayer)
		if hasOverthrown {
			newTurn.FillTurn(throw)
			players.CurrentPlayer.Turns = append(players.CurrentPlayer.Turns, *newTurn)
			x01Engine.CalculatePlayerScore(players.CurrentPlayer)
			players.SwitchToNextPlayer()
			return
		} else {
			newTurn.Append(throw)
			players.CurrentPlayer.Turns = append(players.CurrentPlayer.Turns, *newTurn)
			x01Engine.CalculatePlayerScore(players.CurrentPlayer)
			return
		}
	}
	hasOverthrown := x01Engine.checkForOverThrow(throw, players.CurrentPlayer)
	if hasOverthrown {
		latestTurn.FillTurn(throw)
		x01Engine.CalculatePlayerScore(players.CurrentPlayer)
		players.SwitchToNextPlayer()
		return
	} else {
		hasToSwitchPlayer := latestTurn.Append(throw)
		x01Engine.CalculatePlayerScore(players.CurrentPlayer)
		if hasToSwitchPlayer {
			players.SwitchToNextPlayer()
			return
		}
	}
}

func (x01Engine *X01Engine) UndoLastThrow(players *engine.Players) {
	latestTurnIndex := len(players.CurrentPlayer.Turns) - 1
	latestTurn := &players.CurrentPlayer.Turns[latestTurnIndex]
	if latestTurn.First == nil {
		players.SwitchToPreviousPlayer()
		x01Engine.UndoLastThrow(players)
	}

	if latestTurn.Third != nil {
		latestTurn.Third = nil
	} else if latestTurn.Second != nil {
		latestTurn.Second = nil
	} else if latestTurn.First != nil {
		latestTurn.First = nil
	}
}

func (x01Engine *X01Engine) CalculatePlayerScore(player *engine.Player) {
	var totalSum int16
	for _, turn := range player.Turns {
		totalSum += turn.Sum()
	}
	player.Value.Score = x01Engine.StartingScore - totalSum
}

func (x01Engine *X01Engine) HasAnyPlayerWon(players *engine.Players) *engine.Player {
	head := players.Head
	for head != nil {
		if head.Value.Score == 0 {
			return head
		}
		head = head.Next
	}
	return nil
}

func (x01Engine *X01Engine) checkForOverThrow(throw *domain.Throw, player *engine.Player) bool {
	if throw.Points == 0 {
		return false
	}
	var tempScore int16
	for _, turn := range player.Turns {
		tempScore += turn.Sum()
	}
	tempScore += throw.Points * throw.Multiplicator
	if (x01Engine.StartingScore - tempScore) < 0 {
		logger.Trace("trace: overthrown: %d", x01Engine.StartingScore-tempScore)
		return true
	}
	return false
}

// New returns a new instance of the x01 engine with given starting score and possible points
func New(startingScore int16) engine.Engine {
	return &X01Engine{StartingScore: startingScore, Points: []int16{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 29, 20, 25, 50,
	}}
}
