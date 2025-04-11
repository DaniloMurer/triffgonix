package socket

type MessageType string

const (

	// TODO: implement event types the client can send
	Throw     MessageType = "throw"
	Handshake MessageType = "handshake"
	Save      MessageType = "save"
	UndoThrow MessageType = "undo-throw"
	NewGame   MessageType = "new-game"
	GameState MessageType = "game-state"
	Games     MessageType = "games"
)

type IncomingMessage struct {
	Type    *MessageType   `json:"type"`
	Content map[string]any `json:"content"`
}

type ThrowMessage struct {
	Points        int16 `json:"points"`
	Multiplicator int16 `json:"multiplicator"`
}

type UndoThrowMessage struct{}

type OutgoingMessage struct {
	Type    MessageType `json:"type"`
	Content interface{} `json:"content"`
}
