package engine

import (
	"server/core/domain"
)

type X01Engine struct{}

func (engine X01Engine) GetPlayerThrows(player *domain.Player, turns []Turn) *[]domain.Throw {
	var playerThrows []domain.Throw
	for _, turn := range turns {
		if turn.PlayerId == player.Id {
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
	}
	return &playerThrows
}

func (engine *X01Engine) NextPlayer(players *Players) *domain.Player {
	return players.NextPlayer().Value
}

func (engine *X01Engine) RegisterThrow(throw *domain.Throw, turns *[]Turn, players *Players) {
	// FIXME: it would make more sense to store the turns in the player directly
	// so that we can easily access only his turns instead to manunally filter them
	// TODO: this pointer porn here needs to be fixed because what in the actual fuck is even this
	latestTurnIndex := len(*turns) - 1
	if latestTurnIndex < 0 {
		newTurn := &Turn{PlayerId: players.CurrentPlayer.Value.Id}
		newTurn.Append(throw)
		(*turns) = append((*turns), *newTurn)
		return
	}
	latestTurn := &(*turns)[latestTurnIndex]
	needsNewTurn := latestTurn.Append(throw)
	if needsNewTurn {
		newTurn := &Turn{PlayerId: players.CurrentPlayer.Value.Id}
		newTurn.Append(throw)
		(*turns)[latestTurnIndex] = *newTurn
		players.NextPlayer()
	}
	(*turns)[latestTurnIndex] = *latestTurn
}

func (engine *X01Engine) UndoThrow(throw *domain.Throw, turns *[]Turn, players *Players) {
	latestTurnIndex := len(*turns) - 1
	latestTurn := (*turns)[latestTurnIndex]
	if latestTurn.Third != nil {
		latestTurn.Third = nil
	} else if latestTurn.Second != nil {
		latestTurn.Second = nil
	} else if latestTurn.First != nil {
		latestTurn.First = nil
		players.PreviousPlayer()
	}
}
