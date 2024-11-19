package domain

import "server/internal/triffgonix/models"

type Player struct {
	Id            uint   `json:"id"`
	PlayerName    string `json:"playerName"`
	Score         int16  `json:"score"`
	AveragePoints int16  `json:"averagePoints"`
}

func (player *Player) ToPlayerEntity() *models.Player {
	return &models.Player{
		Model:      models.Model{Id: player.Id},
		PlayerName: player.PlayerName,
	}
}

func (player *Player) FromPlayerEntity(playerEntity *Player) *Player {
	return &Player{
		Id:         playerEntity.Id,
		PlayerName: playerEntity.PlayerName,
	}
}

type Game struct {
	Id      uint     `json:"id"`
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}

func (game *Game) ToGameEntity() *models.Game {
	var players []models.Player
	for _, player := range game.Players {
		players = append(players, *player.ToPlayerEntity())
	}
	return &models.Game{
		Model:   models.Model{Id: uint(game.Id)},
		Name:    game.Name,
		Players: players,
	}
}

func (game *Game) FromGameEntity(gameEntity *Game) *Game {
	var players []Player
	for _, player := range gameEntity.Players {
		newPlayer := Player{}
		newPlayer.FromPlayerEntity(&player)
		players = append(players, newPlayer)
	}

	return &Game{
		Id:      gameEntity.Id,
		Name:    gameEntity.Name,
		Players: players,
	}
}

type Throw struct {
	Id            uint  `json:"id"`
	Points        int16 `json:"points"`
	Multiplicator int16 `json:"multiplicator"`
	PlayerId      uint  `json:"playerId"`
}

func (throw *Throw) ToThrowEntity() *models.Throw {
	return &models.Throw{
		Model:         models.Model{Id: throw.Id},
		Points:        throw.Points,
		Multiplicator: throw.Multiplicator,
		PlayerId:      throw.PlayerId,
	}
}

func (throw *Throw) FromThrowEntity(throwEntity *Throw) *Throw {
	return &Throw{
		Id:            throwEntity.Id,
		Points:        throwEntity.Points,
		Multiplicator: throwEntity.Multiplicator,
		PlayerId:      throwEntity.PlayerId,
	}
}
