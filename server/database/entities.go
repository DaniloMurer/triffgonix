package database

import (
	"time"

	"gorm.io/gorm"
)

// Model base struct for database models
type Model struct {
	Id        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Player struct {
	Model
	PlayerName string `gorm:"unique" json:"username"`
	// TODO: make player - throws connection
}

type Throw struct {
	Model
	Points        int16 `json:"points"`
	Multiplicator int16 `json:"multiplicator"`
	PlayerId      uint  `json:"playerId"`
}

type Game struct {
	Model
	Name    string   `json:"name"`
	Players []Player `gorm:"many2many:game_players;"`
}
