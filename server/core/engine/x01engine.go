package engine

import (
	"server/core/domain"
)

type X01Engine struct {
	StartingScore uint16
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

func (engine *X01Engine) NextPlayer(players *Players) *domain.Player {
	return players.SwitchToNextPlayer().Value
}

func (engine *X01Engine) RegisterThrow(throw *domain.Throw, players *Players) {
	latestTurnIndex := len(players.CurrentPlayer.Turns) - 1
	// if player has no turns, then one should be created first
	if latestTurnIndex < 0 {
		newTurn := &Turn{}
		newTurn.Append(throw)
		players.CurrentPlayer.Turns = append(players.CurrentPlayer.Turns, *newTurn)
		engine.CalculatePlayerScore(players.CurrentPlayer)
		return
	}
	latestTurn := &players.CurrentPlayer.Turns[latestTurnIndex]
	if !latestTurn.HasSpace() {
		newTurn := &Turn{}
		newTurn.Append(throw)
		players.CurrentPlayer.Turns = append(players.CurrentPlayer.Turns, *newTurn)
		engine.CalculatePlayerScore(players.CurrentPlayer)
		return
	}
	hasToSwitchPlayer := latestTurn.Append(throw)
	engine.CalculatePlayerScore(players.CurrentPlayer)
	if hasToSwitchPlayer {
		players.SwitchToNextPlayer()
		return
	}
}

func (engine *X01Engine) UndoThrow(throw *domain.Throw, players *Players) {
	latestTurnIndex := len(players.CurrentPlayer.Turns) - 1
	latestTurn := players.CurrentPlayer.Turns[latestTurnIndex]
	if latestTurn.Third != nil {
		latestTurn.Third = nil
	} else if latestTurn.Second != nil {
		latestTurn.Second = nil
	} else if latestTurn.First != nil {
		latestTurn.First = nil
		players.SwitchToPreviousPlayer()
	}
}

func (engine *X01Engine) CalculatePlayerScore(player *Player) {
	var totalSum uint16
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
