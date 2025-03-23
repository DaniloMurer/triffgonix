package apiplayer

import (
	"github.com/DaniloMurer/triffgonix/server/internal/api/dto"
	"github.com/DaniloMurer/triffgonix/server/internal/database"
	"github.com/DaniloMurer/triffgonix/server/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

var logger = logging.NewLogger()

// CreatePlayer godoc
// @Summary Create a new player
// @Description Creates a new player in the system
// @Tags players
// @Accept json
// @Produce json
// @Param player body dto.Player true "Player information"
// @Success 201 {object} models.Player "Created player"
// @Failure 500 "Internal Server Error"
// @Router /api/user [post]
func CreatePlayer(c *gin.Context) {
	var player dto.Player
	err := c.BindJSON(&player)
	if err != nil {
		logger.Error("error while parsing player json")
		c.Status(http.StatusInternalServerError)
		return
	}
	err, newPlayer := database.CreatePlayer(player.ToEntity())
	if err != nil {
		logger.Error("error while saving player to database: %+v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, &newPlayer)
}

// GetPlayers godoc
// @Summary Get all players
// @Description Retrieves all players from the system
// @Tags players
// @Produce json
// @Success 200 {array} dto.Player "List of players"
// @Router /api/user [get]
func GetPlayers(c *gin.Context) {
	users := database.FindAllUsers()
	var userDtos []dto.Player
	for _, user := range users {
		userDto := dto.Player{}
		userDtos = append(userDtos, userDto.FromEntity(&user))
	}
	c.JSON(http.StatusOK, &userDtos)
}
