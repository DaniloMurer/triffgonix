package apigame

import (
	"github.com/DaniloMurer/triffgonix/server/internal/api"
	"github.com/DaniloMurer/triffgonix/server/internal/api/dto"
	"github.com/DaniloMurer/triffgonix/server/internal/dart/engine"
	"github.com/DaniloMurer/triffgonix/server/internal/dart/engine/x01"
	"github.com/DaniloMurer/triffgonix/server/internal/database"
	"github.com/DaniloMurer/triffgonix/server/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

var logger = logging.NewLogger()

// CreateGame godoc
// @Summary Create a new game
// @Description Creates a new dart game and sets up the corresponding hub
// @Tags games
// @Accept json
// @Produce json
// @Param game body dto.Game true "Game information"
// @Success 201 {object} models.Game "Created game"
// @Failure 500 "Internal Server Error"
// @Router /api/game [post]
func CreateGame(c *gin.Context) {
	var newGame dto.Game
	err := c.BindJSON(&newGame)
	if err != nil {
		logger.Error("error while binding json: %+v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	logger.Trace("new game to be saved: %+v", &newGame)
	err, savedGame := database.CreateGame(newGame.ToEntity())
	if err != nil {
		logger.Error("error while saving game to the databse: %+v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// create new game and hub
	players := engine.Players{}
	for _, player := range newGame.Players {
		players.Add(&engine.Player{Value: player.ToDomain()})
	}
	game := engine.Game{
		Name:    newGame.Name,
		Players: &players,
		Engine:  x01.New(newGame.StartingScore),
	}
	api.CreateHub(savedGame, game)
	c.JSON(http.StatusCreated, &savedGame)
}

// GetGames godoc
// @Summary Get all games
// @Description Retrieves all games from the system
// @Tags games
// @Produce json
// @Success 200 {array} models.Game "List of games"
// @Router /api/game [get]
func GetGames(c *gin.Context) {
	games := database.FindAllGames()
	c.JSON(http.StatusOK, &games)
}
