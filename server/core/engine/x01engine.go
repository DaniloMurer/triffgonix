package engine

import (
	"server/core/domain"
)

type X01Engine struct{}

func (engine X01Engine) GetPlayerThrows(player *Player) *[]domain.Throw {
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
	return players.NextPlayer().Value
}

func (engine *X01Engine) RegisterThrow(throw *domain.Throw, players *Players) {
	latestTurnIndex := len(players.CurrentPlayer.Turns) - 1
	if latestTurnIndex < 0 {
		newTurn := &Turn{}
		newTurn.Append(throw)
		players.CurrentPlayer.Turns = append(players.CurrentPlayer.Turns, *newTurn)
		return
	}
	latestTurn := &players.CurrentPlayer.Turns[latestTurnIndex]
	needsNewTurn := latestTurn.Append(throw)
	if needsNewTurn {
		newTurn := &Turn{}
		newTurn.Append(throw)
		players.CurrentPlayer.Turns[latestTurnIndex] = *newTurn
		players.NextPlayer()
		return
	}
	players.CurrentPlayer.Turns[latestTurnIndex] = *latestTurn
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
		players.PreviousPlayer()
	}
}
