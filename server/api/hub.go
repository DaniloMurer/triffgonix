package handlers

import (
	"server/api/dto"
	"server/core/domain"
	"server/dart/engine"

	"github.com/gorilla/websocket"
)

type Client struct {
	Id         string
	Connection *websocket.Conn
}

type Hub struct {
	Id      uint
	Clients map[*Client]bool
	Game    engine.Game
}

// RegisterNewClient adds a web socket connection to the hub
func (hub *Hub) RegisterNewClient(conn *websocket.Conn) {
	logger.Trace("new client connected")
	client := &Client{Id: "test", Connection: conn}
	hub.Clients[client] = true
}

// broadcastMessage sends given message to all clients connected to the hub
func (hub *Hub) broadcastMessage(message dto.Players) {
	logger.Trace("message to broadcast: %+v", message)
	for client := range hub.Clients {
		client.Connection.WriteJSON(message)
	}
}

func (hub *Hub) HandleConnection(conn *websocket.Conn) {
	defer conn.Close()
	for {
		var message dto.Message
		err := conn.ReadJSON(&message)
		if err != nil {
			logger.Warn("error occured while reading json: %+v", err)
			conn.Close()
			hub.cleanupClient(conn)
			return
		}
		switch *message.Type {
		case dto.Throw:
			// converting with type assertion. actually quite fancy
			points, pointsOk := message.Content["points"].(int16)
			muliplicator, multiplicatorOk := message.Content["multiplicator"].(int16)

			if pointsOk && multiplicatorOk {
				hub.Game.Engine.RegisterThrow(&domain.Throw{Points: points, Multiplicator: muliplicator}, hub.Game.Players)
				hub.broadcastMessage(hub.Game.Players.ToDto())
			} else {
				logger.Warn("couldn't parse throw event's points and multiplicator: %+v", message)
			}
		case dto.UndoThrow:
			hub.Game.Engine.UndoLastThrow(hub.Game.Players)
			hub.broadcastMessage(hub.Game.Players.ToDto())
		default:
			logger.Trace("some other message type received: %s", *message.Type)
		}
	}
}

// cleanupClient removes connection from the hub after disconnect
func (hub *Hub) cleanupClient(conn *websocket.Conn) {
	for client := range hub.Clients {
		if client.Connection == conn {
			delete(hub.Clients, client)
		}
	}
}
