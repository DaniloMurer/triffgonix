package handlers

import (
  "net/http"
  "server/internal/triffgonix/api/dto"
  "server/internal/triffgonix/dart/engine"
  "server/internal/triffgonix/dart/engine/x01"
  "server/internal/triffgonix/database"
  "server/internal/triffgonix/models"
  "server/pkg/logging"
  "strconv"

  "github.com/gin-gonic/gin"
  "github.com/gorilla/websocket"
)

var (
  upgrader           = websocket.Upgrader{}
  hubs               = map[string]Hub{}
  generalConnections []*websocket.Conn
)

var logger logging.Logger = logging.NewLogger()

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
  var message dto.IncomingMessage
  err = conn.ReadJSON(&message)
  if err != nil {
    logger.Error("error while reading from socket connection: %v", err)
    return
  }
  switch *message.Type {
  case dto.Handshake:
    hub, exists := hubs[gameId]
    if exists {
      hub.RegisterNewClient(conn)
    }
  }
  hub, exists := hubs[gameId]
  if exists {
    go hub.HandleConnection(conn)
  } else {
    conn.Close()
  }
}

func HandleGeneralWebsocket(c *gin.Context) {
  upgrader.CheckOrigin = func(r *http.Request) bool {
    return true
  }
  conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
  if err != nil {
    logger.Error("error while upgrading request to websocket protocol: %v", err)
    return
  }
  generalConnections = append(generalConnections, conn)
}

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

func GetPlayers(c *gin.Context) {
  users := database.FindAllUsers()
  c.JSON(http.StatusFound, &users)
}

func CreateGame(c *gin.Context) {
  var newGame dto.Game
  err := c.BindJSON(&newGame)
  if err != nil {
    logger.Error("error while binding json: %+v", err)
    c.Status(http.StatusInternalServerError)
    return
  }
  logger.Trace("new game to be saved: %+v", &newGame)
  err, savedGame := database.CreateGame(newGame.ToEntity())
  if err != nil {
    logger.Error("error while saving game to the databse: %+v", err)
    c.Status(http.StatusInternalServerError)
    return
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
  newHub := Hub{Id: savedGame.Id, Clients: map[*Client]bool{}, Game: game}
  hubs[strconv.FormatUint(uint64(savedGame.Id), 10)] = newHub
  broadcastNewGame(savedGame)
  c.JSON(http.StatusCreated, &savedGame)
}

func GetGames(c *gin.Context) {
  games := database.FindAllGames()
  c.JSON(http.StatusFound, &games)
}

func broadcastNewGame(newGame *models.Game) {
  game := dto.Game{}
  game.FromEntity(newGame)
  message := dto.OutgoingMessage{Type: dto.NewGame, Content: game}
  for _, conn := range generalConnections {
    err := conn.WriteJSON(message)
    if err != nil {
      logger.Error("error while writing new game json: %+v", err)
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
