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

// HandleDartWebSocket Manages the connection lifecycle for a dart game session.
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
	// FIXME: review this code. seems to be weird. i add the connection to the hub if handshake
	// to then later handle the connection on a separate thread. maybe i can merge this? did i
	// do this as a sort of rejoin feature? doesn't seem to be thought through tho
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

// HandleGeneralWebsocket Sends a list of games to the client and tracks the connection,
// and handles connections.
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
		// FIXME: multiple connections don't work here. broken pipe. question is, if that's an issue
		// when calling from the same browser instance, if not we have a larger issue
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
