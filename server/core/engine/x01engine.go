package engine

import "server/core/domain"

type X01Engine struct {
	Game Game
}

func (engine X01Engine) getPlayerThrows(player *domain.Player) *[]domain.Throw {
	// TODO: write implementation
	return nil
}

func (engine X01Engine) getNextPlayer(player *Player) *domain.Player {
	// TODO: write implementation
	return nil
}

func (engine X01Engine) RegisterThrow(throw *domain.Throw) uint {
	// TODO: write implementation
	return 0
}

func (engine X01Engine) UndoThrow(throw *domain.Throw) uint {
	// TODO: write implementation
	return 0
}
