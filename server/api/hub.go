package handlers

import (
	"log"
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

// RegsiterNewClient adds a web socket connection to the hub
func (hub *Hub) RegsiterNewClient(conn *websocket.Conn) {
	log.Println("trace: new client connected")
	client := &Client{Id: "test", Connection: conn}
	hub.Clients[client] = true
}

// BroadcastMessage sends given message to all clients connected to the hub
func (hub *Hub) BroadcastMessage(message interface{}) {
	for client := range hub.Clients {
		client.Connection.WriteJSON(message)
	}
}

// TODO: implement function for handling websocket messages in a loop. will be called as goroutine
