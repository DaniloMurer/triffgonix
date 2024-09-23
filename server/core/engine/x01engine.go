package engine

import "server/core/domain"

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

func (engine X01Engine) NextPlayer(players *Players) *domain.Player {
	return players.NextPlayer().Value
}

func (engine X01Engine) RegisterThrow(throw *domain.Throw, turns []Turn) {
	latestTurnIndex := len(turns) - 1
	latestTurn := &turns[latestTurnIndex]
	needsNewTurn := latestTurn.Append(throw)
	if needsNewTurn {
		newTurn := Turn{}
		newTurn.Append(throw)
	}
}

func (engine X01Engine) UndoThrow(throw *domain.Throw, turns []Turn) {
	// FIXME: at the moment there is no way to tell when the previous player should be returned
	// TODO: write implementation
}
