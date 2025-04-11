package apigame

import (
	"github.com/DaniloMurer/triffgonix/server/internal/api"
	"github.com/DaniloMurer/triffgonix/server/internal/api/dto"
	"github.com/DaniloMurer/triffgonix/server/internal/dart/engine"
	"github.com/DaniloMurer/triffgonix/server/internal/dart/engine/x01"
	"github.com/DaniloMurer/triffgonix/server/internal/database"
	"github.com/DaniloMurer/triffgonix/server/pkg/logging"
	"github.com/gofiber/fiber/v2"
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
// @Success 201 {object} dto.Game "Created game"
// @Failure 500 "Internal Server Error"
// @Router /api/game [post]
func CreateGame(c *fiber.Ctx) error {
	var newGame dto.Game
	err := c.BodyParser(&newGame)
	if err != nil {
		logger.Error("error while binding json: %+v", err)
		c.Status(http.StatusInternalServerError)
		return nil
	}
	logger.Trace("new game to be saved: %+v", &newGame)
	err, savedGame := database.CreateGame(newGame.ToEntity())
	if err != nil {
		logger.Error("error while saving game to the databse: %+v", err)
		c.Status(http.StatusInternalServerError)
		return nil
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

	var gameDto dto.Game
	gameDto.FromEntity(savedGame)
	return c.Status(http.StatusCreated).JSON(&gameDto)
}

// GetGames godoc
// @Summary Get all games
// @Description Retrieves all games from the system
// @Tags games
// @Produce json
// @Success 200 {array} models.Game "List of games"
// @Router /api/game [get]
func GetGames(c *fiber.Ctx) error {
	games := database.FindAllGames()
	return c.Status(http.StatusOK).JSON(&games)
}
