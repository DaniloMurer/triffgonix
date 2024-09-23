package engine

import (
	"server/core/domain"
	"testing"
)

// if we're at the first player, the netx player should be the second one in the list
func TestGetNextPlayerWhenFirstPlayerTurn(t *testing.T) {
	player := Player{Value: &domain.Player{PlayerName: "1"}, Previous: nil, Next: nil}
	player2 := Player{Value: &domain.Player{PlayerName: "2"}, Previous: &player, Next: nil}
	player.Next = &player2

	players := Players{Head: &player, CurrentPlayer: &player, Tail: &player2}

	game := Game{Name: "we", Players: players, Engine: X01Engine{}}

	nextPlayer := game.Engine.NextPlayer(&game.Players)
	if nextPlayer.PlayerName != "2" {
		t.Fatalf(`ERROR: The next player should have been player 2. Instead got %q`, nextPlayer.PlayerName)
	}
}

// if we're at the last player, the next player should be the first in the list
func TestGetNextPlayerWhenLastPlayerTurn(t *testing.T) {
	player := Player{Value: &domain.Player{PlayerName: "1"}, Previous: nil, Next: nil}
	player2 := Player{Value: &domain.Player{PlayerName: "2"}, Previous: &player, Next: nil}
	player.Next = &player2

	players := Players{Head: &player, CurrentPlayer: &player2, Tail: &player2}

	game := Game{Name: "we", Players: players, Engine: X01Engine{}}

	nextPlayer := game.Engine.NextPlayer(&game.Players)
	if nextPlayer.PlayerName != "1" {
		t.Fatalf(`ERROR: The next player should have been player 1. Instead got %q`, nextPlayer.PlayerName)
	}
}
