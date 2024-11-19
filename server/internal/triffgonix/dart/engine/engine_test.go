package engine

import (
	"server/internal/triffgonix/domain"
	"testing"
)

// Mock implementations for interfaces or methods not provided in the original code

// Test the GetAveragePoints method
func TestGetAveragePoints(t *testing.T) {
	player := &Player{
		Turns: []Turn{
			{First: &domain.Throw{Points: 10, Multiplicator: 1}, Second: &domain.Throw{Points: 20, Multiplicator: 1}},
			{First: &domain.Throw{Points: 15, Multiplicator: 2}, Second: &domain.Throw{Points: 0, Multiplicator: 1}},
		},
	}

	averagePoints := player.GetAveragePoints()
	expectedPoints := int16(60 / 4)

	if averagePoints != expectedPoints {
		t.Errorf("expected: %d, got: %d", expectedPoints, averagePoints)
	}
}

// Test the Add method in Players
func TestPlayers_Add(t *testing.T) {
	players := &Players{}
	player1 := &Player{Value: &domain.Player{PlayerName: "Player 1"}}
	player2 := &Player{Value: &domain.Player{PlayerName: "Player 2"}}

	players.Add(player1)
	players.Add(player2)

	if players.Head != player1 {
		t.Errorf("expected head to be player1, got: %v", players.Head)
	}
	if players.Tail != player2 {
		t.Errorf("expected tail to be player2, got: %v", players.Tail)
	}
	if player1.Next != player2 {
		t.Errorf("expected player1 to be followed by player2, got: %v", player1.Next)
	}
}

// Test the SwitchToNextPlayer method
func TestPlayers_SwitchToNextPlayer(t *testing.T) {
	players := &Players{}
	player1 := &Player{Value: &domain.Player{PlayerName: "Player 1"}}
	player2 := &Player{Value: &domain.Player{PlayerName: "Player 2"}}

	players.Add(player1)
	players.Add(player2)

	players.SwitchToNextPlayer()

	if players.CurrentPlayer != player2 {
		t.Errorf("expected current player to be player2, got: %v", players.CurrentPlayer)
	}
}

// Test the Sum method in Turn
func TestTurn_Sum(t *testing.T) {
	turn := &Turn{
		First:  &domain.Throw{Points: 10, Multiplicator: 2},
		Second: &domain.Throw{Points: 5, Multiplicator: 1},
		Third:  &domain.Throw{Points: 20, Multiplicator: 3},
	}

	total := turn.Sum()
	expectedTotal := int16(10*2 + 5*1 + 20*3)

	if total != expectedTotal {
		t.Errorf("expected sum to be %d, got: %d", expectedTotal, total)
	}
}

// Test the Append method in Turn
func TestTurn_Append(t *testing.T) {
	turn := &Turn{}

	throw1 := &domain.Throw{Points: 10, Multiplicator: 2}
	throw2 := &domain.Throw{Points: 5, Multiplicator: 1}
	throw3 := &domain.Throw{Points: 20, Multiplicator: 3}

	if turn.Append(throw1) {
		t.Errorf("expected false as the turn isn't full, got: true")
	}
	if turn.Append(throw2) {
		t.Errorf("expected false as the turn still has space, got: true")
	}
	if !turn.Append(throw3) {
		t.Errorf("expected true as the turn becomes full, got: false")
	}
	if turn.First != throw1 || turn.Second != throw2 || turn.Third != throw3 {
		t.Errorf("throws not appended correctly")
	}
}

// Test the HasSpace method in Turn
func TestTurn_HasSpace(t *testing.T) {
	turn := &Turn{}

	if !turn.HasSpace() {
		t.Errorf("expected true initially, got: false")
	}

	turn.First = &domain.Throw{}
	turn.Second = &domain.Throw{}
	turn.Third = &domain.Throw{}

	if turn.HasSpace() {
		t.Errorf("expected false when all throws are set, got: true")
	}
}

// Test the FillTurn method in Turn
func TestTurn_FillTurn(t *testing.T) {
	turn := &Turn{}
	throw := &domain.Throw{PlayerId: 1}

	turn.FillTurn(throw)

	if turn.First.Points != 0 || turn.Second.Points != 0 || turn.Third.Points != 0 {
		t.Errorf("expected throws to be filled with 0 points")
	}
	if turn.First.PlayerId != throw.PlayerId || turn.Second.PlayerId != throw.PlayerId || turn.Third.PlayerId != throw.PlayerId {
		t.Errorf("expected player IDs to match")
	}
}

// Test the ThrowCount method in Turn
func TestTurn_ThrowCount(t *testing.T) {
	turn := &Turn{}

	if count := turn.ThrowCount(); count != 0 {
		t.Errorf("expected 0 count initially, got: %d", count)
	}

	turn.First = &domain.Throw{}
	if count := turn.ThrowCount(); count != 1 {
		t.Errorf("expected 1, got: %d", count)
	}

	turn.Second = &domain.Throw{}
	if count := turn.ThrowCount(); count != 2 {
		t.Errorf("expected 2, got: %d", count)
	}

	turn.Third = &domain.Throw{}
	if count := turn.ThrowCount(); count != 3 {
		t.Errorf("expected 3, got: %d", count)
	}
}

// Test the ToDto method in Players
func TestPlayers_ToDto(t *testing.T) {
	player1 := &Player{Value: &domain.Player{PlayerName: "Player 1"}}
	player2 := &Player{Value: &domain.Player{PlayerName: "Player 2"}}
	players := &Players{}
	players.Add(player1)
	players.Add(player2)
	players.CurrentPlayer = player2

	dtoPlayers := players.ToDto()

	if len(dtoPlayers.AllPlayers) != 2 {
		t.Errorf("expected 2 players in dto, got: %d", len(dtoPlayers.AllPlayers))
	}
	if dtoPlayers.CurrentPlayer.PlayerName != player2.Value.PlayerName {
		t.Errorf("expected current player to be Player 2, got: %s", dtoPlayers.CurrentPlayer.PlayerName)
	}
}
