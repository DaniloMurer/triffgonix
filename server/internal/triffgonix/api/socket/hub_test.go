package socket

import (
	"server/pkg/logging"
	"testing"
)

// Mock WebSocket connection
type MockConn struct {
	WriteJSONFunc func(v interface{}) error
	ReadJSONFunc  func(v interface{}) error
	CloseFunc     func() error
}

func (m *MockConn) WriteJSON(v interface{}) error {
	return m.WriteJSONFunc(v)
}

func (m *MockConn) ReadJSON(v interface{}) error {
	return m.ReadJSONFunc(v)
}

func (m *MockConn) Close() error {
	return m.CloseFunc()
}

func TestRegisterNewClient(t *testing.T) {
	hub := &Hub{Clients: make(map[*Client]bool)}
	mockConn := &MockConn{}
	logger = logging.NewLogger() // Ensure logger is initialized

	hub.RegisterNewClient(mockConn)

	if len(hub.Clients) != 1 {
		t.Errorf("expected 1 client in hub, got: %d", len(hub.Clients))
	}
}

func TestBroadcastMessage(t *testing.T) {
	hub := &Hub{Clients: make(map[*Client]bool)}
	mockConn := &MockConn{
		WriteJSONFunc: func(v interface{}) error {
			return nil
		},
	}
	client := &Client{Connection: mockConn}
	hub.Clients[client] = true
	logger = logging.NewLogger() // Ensure logger is initialized

	message := OutgoingMessage{Type: GameState, Content: "test content"}
	hub.broadcastMessage(message)
	// No assertion needed. If WriteJSON panics, the test will fail
}

func TestCleanupClient(t *testing.T) {
	hub := &Hub{Clients: make(map[*Client]bool)}
	mockConn := &MockConn{}
	client := &Client{Connection: mockConn}
	hub.Clients[client] = true
	logger = logging.NewLogger() // Ensure logger is initialized

	if len(hub.Clients) != 1 {
		t.Errorf("expected 1 client, got: %d", len(hub.Clients))
	}

	hub.cleanupClient(mockConn)

	if len(hub.Clients) != 0 {
		t.Errorf("expected 0 clients after cleanup, got: %d", len(hub.Clients))
	}
}
