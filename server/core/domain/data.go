package domain

import "server/database"

type Player struct {
	Id         uint   `json:"id"`
	PlayerName string `json:"playerName"`
}

func (self *Player) ToPlayerEntity() *database.Player {
	return &database.Player{
		Model:      database.Model{Id: self.Id},
		PlayerName: self.PlayerName,
	}
}

func (self *Player) FromPlayerEntity(player *database.Player) *Player {
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

func (self *Game) ToGameEntity() *database.Game {
	var players []database.Player
	for _, player := range self.Players {
		players = append(players, *player.ToPlayerEntity())
	}
	return &database.Game{
		Model:   database.Model{Id: uint(self.Id)},
		Name:    self.Name,
		Players: players,
	}
}

func (self *Game) FromGameEntity(game *database.Game) *Game {
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
	Points        uint8 `json:"points"`
	Multiplicator uint8 `json:"multiplicator"`
	PlayerId      uint  `json:"playerId"`
}

func (self *Throw) ToThrowEntity() *database.Throw {
	return &database.Throw{
		Model:         database.Model{Id: self.Id},
		Points:        self.Points,
		Multiplicator: self.Multiplicator,
		PlayerId:      self.PlayerId,
	}
}

func (self *Throw) FromThrowEntity(throw *database.Throw) *Throw {
	return &Throw{
		Id:            throw.Id,
		Points:        throw.Points,
		Multiplicator: throw.Multiplicator,
		PlayerId:      throw.PlayerId,
	}
}
