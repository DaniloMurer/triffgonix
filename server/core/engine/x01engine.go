package engine

import "server/core/domain"

type X01Engine struct {
	Game Game
}

func (engine X01Engine) getPlayerThrows(player *domain.Player) *[]domain.Throw {
	var throws []domain.Throw
	for _, throw := range engine.Game.Throws {
		if throw.PlayerId == player.Id {
			throws = append(throws, throw)
		}
	}
	return &throws
}

func (engine X01Engine) nextPlayer() *domain.Player {
	return engine.Game.Players.NextPlayer().Value
}

func (engine X01Engine) RegisterThrow(throw *domain.Throw) uint {
	// FIXME: at the moment there is no real way to tell when the player should switch after a throw
	// TODO: write implementation
	return 0
}

func (engine X01Engine) UndoThrow(throw *domain.Throw) uint {
	// FIXME: at the moment there is no way to tell when the previous player should be returned
	// TODO: write implementation
	return 0
}
