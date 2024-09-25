package engine

import (
	"fmt"
	"server/core/domain"
)

type X01Engine struct {
	StartingScore int16
}

func (engine *X01Engine) GetPlayerThrows(player *Player) *[]domain.Throw {
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

func (engine *X01Engine) RegisterThrow(throw *domain.Throw, players *Players) {
	latestTurnIndex := len(players.CurrentPlayer.Turns) - 1
	// if player has no turns, then one should be created first
	if latestTurnIndex < 0 {
		newTurn := &Turn{}
		hasOverthrown := engine.checkForOverThrow(throw, players.CurrentPlayer)
		if hasOverthrown {
			newTurn.FillTurn(throw)
			players.CurrentPlayer.Turns = append(players.CurrentPlayer.Turns, *newTurn)
			players.SwitchToNextPlayer()
			return
		} else {
			newTurn.Append(throw)
			players.CurrentPlayer.Turns = append(players.CurrentPlayer.Turns, *newTurn)
			engine.CalculatePlayerScore(players.CurrentPlayer)
			return
		}
	}
	latestTurn := &players.CurrentPlayer.Turns[latestTurnIndex]
	if !latestTurn.HasSpace() {
		newTurn := &Turn{}
		hasOverthrown := engine.checkForOverThrow(throw, players.CurrentPlayer)
		if hasOverthrown {
			latestTurn.FillTurn(throw)
			players.CurrentPlayer.Turns = append(players.CurrentPlayer.Turns, *newTurn)
			players.SwitchToNextPlayer()
			return
		} else {
			newTurn.Append(throw)
			players.CurrentPlayer.Turns = append(players.CurrentPlayer.Turns, *newTurn)
			engine.CalculatePlayerScore(players.CurrentPlayer)
			return
		}
	}
	hasOverthrown := engine.checkForOverThrow(throw, players.CurrentPlayer)
	if hasOverthrown {
		latestTurn.FillTurn(throw)
		players.SwitchToNextPlayer()
		return
	} else {
		hasToSwitchPlayer := latestTurn.Append(throw)
		engine.CalculatePlayerScore(players.CurrentPlayer)
		if hasToSwitchPlayer {
			players.SwitchToNextPlayer()
			return
		}
	}
}

func (engine *X01Engine) UndoLastThrow(players *Players) {
	latestTurnIndex := len(players.CurrentPlayer.Turns) - 1
	latestTurn := &players.CurrentPlayer.Turns[latestTurnIndex]
	if latestTurn.First == nil {
		players.SwitchToPreviousPlayer()
		engine.UndoLastThrow(players)
	}

	if latestTurn.Third != nil {
		latestTurn.Third = nil
	} else if latestTurn.Second != nil {
		latestTurn.Second = nil
	} else if latestTurn.First != nil {
		latestTurn.First = nil
	}
}

func (engine *X01Engine) CalculatePlayerScore(player *Player) {
	var totalSum int16
	for _, turn := range player.Turns {
		totalSum += turn.Sum()
	}
	player.Value.Score = engine.StartingScore - totalSum
}

func (engine *X01Engine) HasAnyPlayerWon(players *Players) *Player {
	head := players.Head
	for head != nil {
		if head.Value.Score == 0 {
			return head
		}
		head = head.Next
	}
	return nil
}

func (engine *X01Engine) checkForOverThrow(throw *domain.Throw, player *Player) bool {
	if throw.Points == 0 {
		return false
	}
	var tempScore int16
	for _, turn := range player.Turns {
		tempScore += turn.Sum()
	}
	tempScore += throw.Multiplicator * throw.Points
	if (engine.StartingScore - tempScore) < 0 {
		fmt.Printf("trace: overthrown with score: %d\n", engine.StartingScore-tempScore)
		return true
	}
	return false
}
