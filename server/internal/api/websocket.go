package api

import (
	"github.com/DaniloMurer/triffgonix/server/internal/api/dto"
	"github.com/DaniloMurer/triffgonix/server/internal/api/socket"
	"github.com/DaniloMurer/triffgonix/server/internal/dart/engine"
	"github.com/DaniloMurer/triffgonix/server/internal/models"
	"github.com/DaniloMurer/triffgonix/server/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var (
	upgrader           = websocket.Upgrader{}
	hubs               = map[string]socket.Hub{}
	generalConnections []*websocket.Conn
)

var logger = logging.NewLogger()

func CreateHub(savedGame *models.Game, game engine.Game) {
	newHub := socket.Hub{Id: savedGame.Id, Clients: map[*socket.Client]bool{}, Game: game}
	hubs[strconv.FormatUint(uint64(savedGame.Id), 10)] = newHub
	broadcastNewGame(savedGame)
}

// HandleDartWebSocket godoc
// @Summary Handle game-specific WebSocket connections
// @Description Handles WebSocket connections for specific dart games
// @Tags websocket
// @Accept json
// @Param gameId path string true "Game ID"
// @Success 101 "Switching Protocols to WebSocket"
// @Failure 400 "Bad Request"
// @Router /ws/dart/{gameId} [get]
func HandleDartWebSocket(c *gin.Context) {
	cleanupHubs()
	// FIXME: only temporary
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("error while upgrading request to websocket protocol: %v", err)
		return
	}
	gameId := c.Param("gameId")
	// get message from socket
	var message socket.IncomingMessage
	err = conn.ReadJSON(&message)
	if err != nil {
		logger.Error("error while reading from socket connection: %v", err)
		return
	}
	switch *message.Type {
	case socket.Handshake:
		hub, exists := hubs[gameId]
		if exists {
			hub.RegisterNewClient(conn)
		}
	}
	hub, exists := hubs[gameId]
	if exists {
		go hub.HandleConnection(conn)
	} else {
		err := conn.Close()
		if err != nil {
			logger.Error("error while closing websocket connection: %+v", err)
		}
	}
}

// HandleGeneralWebsocket godoc
// @Summary Handle general WebSocket connections
// @Description Handles WebSocket connections for general game updates
// @Tags websocket
// @Accept json
// @Success 101 "Switching Protocols to WebSocket"
// @Failure 400 "Bad Request"
// @Router /ws/dart [get]
func HandleGeneralWebsocket(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("error while upgrading request to websocket protocol: %v", err)
		return
	}
	var games []engine.GameDto
	for _, hub := range hubs {
		games = append(games, hub.GetGame())
	}
	generalConnections = append(generalConnections, conn)
	err = conn.WriteJSON(socket.OutgoingMessage{Type: socket.Games, Content: games})
	if err != nil {
		logger.Error("error while writing games message: %+v", err)
		return
	}
}

func broadcastNewGame(newGame *models.Game) {
	game := dto.Game{}
	game.FromEntity(newGame)
	for _, conn := range generalConnections {
		var games []engine.GameDto
		for _, hub := range hubs {
			games = append(games, hub.GetGame())
		}
		generalConnections = append(generalConnections, conn)
		err := conn.WriteJSON(socket.OutgoingMessage{Type: socket.Games, Content: games})
		if err != nil {
			logger.Error("error while writing games message: %+v", err)
			return
		}
	}
}

// cleanupHubs removes hubs with zero clients connected to it
func cleanupHubs() {
	for gameId, hub := range hubs {
		if len(hub.Clients) == 0 {
			delete(hubs, gameId)
		}
	}
}
