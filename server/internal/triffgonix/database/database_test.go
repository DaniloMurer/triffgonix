package database

import (
	"github.com/google/uuid"
	"server/internal/triffgonix/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutoMigrate(t *testing.T) {
	assert.NotPanics(t, AutoMigrate)
}

func TestCreatePlayer(t *testing.T) {
	player := &models.Player{PlayerName: uuid.Must(uuid.NewRandom()).String()}
	err, _ := CreatePlayer(player)
	assert.Nil(t, err)
}

func TestFindAllUsers(t *testing.T) {
	users := FindAllUsers()
	assert.NotNil(t, users)
}

func TestCreateGame(t *testing.T) {
	game := &models.Game{Name: "Test Game"}
	err, _ := CreateGame(game)
	assert.Nil(t, err)
}

func TestFindAllGames(t *testing.T) {
	games := FindAllGames()
	assert.NotNil(t, games)
}
