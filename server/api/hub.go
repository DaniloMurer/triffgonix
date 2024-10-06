package handlers

import (
	"fmt"
	"log"
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
	log.Println("trace: new client connected")
	client := &Client{Id: "test", Connection: conn}
	hub.Clients[client] = true
}

// broadcastMessage sends given message to all clients connected to the hub
func (hub *Hub) broadcastMessage(message interface{}) {
	fmt.Printf("trace: message to broadcast: %+v", message)
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
			log.Printf("error: %+v", err)
		}
		switch *message.Type {
		case dto.Throw:
			fmt.Println("trace: broadcasting message")
			hub.Game.Engine.RegisterThrow(&domain.Throw{Points: 1, Multiplicator: 1}, hub.Game.Players)
			hub.broadcastMessage(&hub.Game.Players)
		default:
			fmt.Println("trace: some other message")
		}
	}
}

// TODO: implement function for handling websocket messages in a loop. will be called as goroutine
