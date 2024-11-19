package main

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_RootRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Welcome")
}

func TestMain_GroupRoutes(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		method   string
		endpoint string
		expected int
	}{
		{"POST", "/api/user", 400}, // Assuming it requires a body and proper setup
		{"GET", "/api/user", 302},
		{"POST", "/api/game", 400}, // Assuming it requires a body and proper setup
		{"GET", "/api/game", 302},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(tt.method, tt.endpoint, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, tt.expected, w.Code)
	}
}

func TestMain_WebSocketRoutes(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		method   string
		endpoint string
		expected int
	}{
		{"GET", "/ws/dart/123", 400}, // Assuming it requires a proper WebSocket setup
		{"GET", "/ws/dart", 400},     // Assuming it requires a proper WebSocket setup
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(tt.method, tt.endpoint, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, tt.expected, w.Code)
	}
}
