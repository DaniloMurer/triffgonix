package dto

type MessageType string

const (
	// TODO: implement event types the client can send
	Throw     MessageType = "throw"
	Handshake MessageType = "handshake"
	Save      MessageType = "save"
	UndoThrow MessageType = "undo-throw"
)

type Message struct {
	Type    *MessageType   `json:"type"`
	Content map[string]any `json:"content"`
}
