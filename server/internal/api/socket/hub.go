package socket

import (
	"github.com/DaniloMurer/triffgonix/server/internal/dart/engine"
	"github.com/DaniloMurer/triffgonix/server/internal/domain"
	"github.com/DaniloMurer/triffgonix/server/pkg/logging"
	"github.com/gorilla/websocket"
)

var logger = logging.NewLogger()

// WebSocketConnection defines the behavior required for a WebSocket connection
type WebSocketConnection interface {
	WriteJSON(v interface{}) error
	ReadJSON(v interface{}) error
	Close() error
}

type Client struct {
	Id         string
	Connection WebSocketConnection
}

type Hub struct {
	Id      uint
	Clients map[*Client]bool
	Game    engine.Game
}

// RegisterNewClient adds a web socket connection to the hub
func (hub *Hub) RegisterNewClient(conn WebSocketConnection) {
	logger.Trace("new client connected")
	client := &Client{Id: "test", Connection: conn}
	hub.Clients[client] = true
}

// broadcastMessage sends given message to all clients connected to the hub
func (hub *Hub) broadcastMessage(message OutgoingMessage) {
	for client := range hub.Clients {
		err := client.Connection.WriteJSON(message)
		if err != nil {
			logger.Error("error while sending message to client: %+v", err)
		}
	}
}

func (hub *Hub) BroadcastToClients(obj OutgoingMessage) []error {
	var errors []error
	for client := range hub.Clients {
		err := client.Connection.WriteJSON(obj)
		if err != nil {
			logger.Error("error while broadcasting to client: %+v", err)
			errors = append(errors, err)
		}
	}
	return errors
}

func (hub *Hub) HandleConnection(conn *websocket.Conn) {
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			logger.Error("error while closing websocket connection: %+v", err)
		}
	}(conn)
	for {
		var message IncomingMessage
		err := conn.ReadJSON(&message)
		if err != nil {
			logger.Warn("error occured while reading json: %+v", err)
			connErr := conn.Close()
			if connErr != nil {
				logger.Error("error while closing websocket connection: %+v", connErr)
			}
			hub.cleanupClient(conn)
			return
		}
		switch *message.Type {
		case Throw:
			// converting with type assertion. actually quite fancy
			points, pointsOk := message.Content["points"].(int16)
			muliplicator, multiplicatorOk := message.Content["multiplicator"].(int16)

			if pointsOk && multiplicatorOk {
				hub.Game.Engine.RegisterThrow(&domain.Throw{Points: points, Multiplicator: muliplicator}, hub.Game.Players)
				gameState := OutgoingMessage{Type: GameState, Content: hub.Game.Players.ToDto()}
				hub.broadcastMessage(gameState)
			} else {
				logger.Warn("couldn't parse throw event's points and multiplicator: %+v", message)
			}
		case UndoThrow:
			hub.Game.Engine.UndoLastThrow(hub.Game.Players)
			gameState := OutgoingMessage{Type: GameState, Content: hub.Game.Players.ToDto()}
			hub.broadcastMessage(gameState)
		default:
			logger.Trace("some other message type received: %s", *message.Type)
		}
	}
}

// cleanupClient removes connection from the hub after disconnect
func (hub *Hub) cleanupClient(conn WebSocketConnection) {
	for client := range hub.Clients {
		if client.Connection == conn {
			delete(hub.Clients, client)
		}
	}
}
