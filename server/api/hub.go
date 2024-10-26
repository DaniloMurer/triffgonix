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
	Id      string
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
			return
		}
		switch *message.Type {
		case dto.Throw:
			hub.Game.Engine.RegisterThrow(&domain.Throw{Points: 1, Multiplicator: 1}, hub.Game.Players)
			hub.broadcastMessage(hub.Game.Players.ToDto())
		case dto.UndoThrow:
			hub.Game.Engine.UndoLastThrow(hub.Game.Players)
			hub.broadcastMessage(hub.Game.Players.ToDto())
		default:
			logger.Trace("some other message type received: %s", *message.Type)
		}
	}
}

// TODO: implement function for handling websocket messages in a loop. will be called as goroutine
