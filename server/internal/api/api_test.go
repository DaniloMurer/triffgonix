package api

import (
	"github.com/DaniloMurer/triffgonix/server/internal/database"
	"github.com/DaniloMurer/triffgonix/server/internal/models"
	"github.com/google/uuid"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePlayer(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	group := r.Group("/api")
	{
		group.POST("/user", CreatePlayer)
		group.GET("/user", GetPlayers)
		group.POST("/game", CreateGame)
		group.GET("/game", GetGames)
	}
	return r
}

func TestDatabaseConnections(t *testing.T) {
	database.AutoMigrate()
	err, _ := database.CreatePlayer(&models.Player{PlayerName: uuid.Must(uuid.NewRandom()).String()})
	assert.Nil(t, err)
	games := database.FindAllGames()
	assert.NotNil(t, games)
}
