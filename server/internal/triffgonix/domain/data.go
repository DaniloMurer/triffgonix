package domain

import "server/internal/triffgonix/models"

type Player struct {
	Id            uint   `json:"id"`
	PlayerName    string `json:"playerName"`
	Score         int16  `json:"score"`
	AveragePoints int16  `json:"averagePoints"`
}

func (self *Player) ToPlayerEntity() *models.Player {
	return &models.Player{
		Model:      models.Model{Id: self.Id},
		PlayerName: self.PlayerName,
	}
}

func (self *Player) FromPlayerEntity(player *Player) *Player {
	return &Player{
		Id:         player.Id,
		PlayerName: player.PlayerName,
	}
}

type Game struct {
	Id      uint     `json:"id"`
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}

func (self *Game) ToGameEntity() *models.Game {
	var players []models.Player
	for _, player := range self.Players {
		players = append(players, *player.ToPlayerEntity())
	}
	return &models.Game{
		Model:   models.Model{Id: uint(self.Id)},
		Name:    self.Name,
		Players: players,
	}
}

func (self *Game) FromGameEntity(game *Game) *Game {
	var players []Player
	for _, player := range game.Players {
		newPlayer := Player{}
		newPlayer.FromPlayerEntity(&player)
		players = append(players, newPlayer)
	}

	return &Game{
		Id:      game.Id,
		Name:    game.Name,
		Players: players,
	}
}

type Throw struct {
	Id            uint  `json:"id"`
	Points        int16 `json:"points"`
	Multiplicator int16 `json:"multiplicator"`
	PlayerId      uint  `json:"playerId"`
}

func (self *Throw) ToThrowEntity() *models.Throw {
	return &models.Throw{
		Model:         models.Model{Id: self.Id},
		Points:        self.Points,
		Multiplicator: self.Multiplicator,
		PlayerId:      self.PlayerId,
	}
}

func (self *Throw) FromThrowEntity(throw *Throw) *Throw {
	return &Throw{
		Id:            throw.Id,
		Points:        throw.Points,
		Multiplicator: throw.Multiplicator,
		PlayerId:      throw.PlayerId,
	}
}
