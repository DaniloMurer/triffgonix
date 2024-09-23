package engine

import "server/core/domain"

type X01Engine struct{}

func (engine X01Engine) GetPlayerThrows(player *domain.Player, throws *[]domain.Throw) *[]domain.Throw {
	var playerThrows []domain.Throw
	for _, throw := range *throws {
		if throw.PlayerId == player.Id {
			playerThrows = append(playerThrows, throw)
		}
	}
	return &playerThrows
}

func (engine X01Engine) NextPlayer(players *Players) *domain.Player {
	return players.NextPlayer().Value
}

func (engine X01Engine) RegisterThrow(throw *domain.Throw, throws *[]domain.Throw) {
	// FIXME: at the moment there is no real way to tell when the player should switch after a throw
	// TODO: write implementation
}

func (engine X01Engine) UndoThrow(throw *domain.Throw, throws *[]domain.Throw) {
	// FIXME: at the moment there is no way to tell when the previous player should be returned
	// TODO: write implementation
}
