package handlers

import (
	"server/api/dto"
	"server/core/domain"
	"server/dart/engine"
	"server/logger"

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

var log logger.Logger = logger.NewLogger()

// RegisterNewClient adds a web socket connection to the hub
func (hub *Hub) RegisterNewClient(conn *websocket.Conn) {
	log.Trace("new client connected")
	client := &Client{Id: "test", Connection: conn}
	hub.Clients[client] = true
}

// broadcastMessage sends given message to all clients connected to the hub
func (hub *Hub) broadcastMessage(message interface{}) {
	log.Trace("message to broadcast: %+v", message)
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
			log.Warn("error occured while reading json: %+v", err)
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
			log.Trace("some other message type received: %s", *message.Type)
		}
	}
}

// TODO: implement function for handling websocket messages in a loop. will be called as goroutine
