package engine

import (
	"server/core/domain"
	"testing"
)

// if we're at the first player, the netx player should be the second one in the list
func TestGetNextPlayerWhenFirstPlayerTurn(t *testing.T) {
	player := Player{Value: &domain.Player{PlayerName: "1"}, Previous: nil, Next: nil, Turns: []Turn{}}
	player2 := Player{Value: &domain.Player{PlayerName: "2"}, Previous: &player, Next: nil, Turns: []Turn{}}
	player.Next = &player2

	players := Players{Head: &player, CurrentPlayer: &player, Tail: &player2}

	game := Game{Players: &players, Engine: &X01Engine{}}

	nextPlayer := game.Engine.NextPlayer(game.Players)
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

	game := Game{Players: &players, Engine: &X01Engine{}}

	nextPlayer := game.Engine.NextPlayer(game.Players)
	if nextPlayer.PlayerName != "1" {
		t.Fatalf(`ERROR: The next player should have been player 1. Instead got %q`, nextPlayer.PlayerName)
	}
}

func TestRegisterThrowToPlayer(t *testing.T) {
	/*player := Player{Value: &domain.Player{Id: 1}, Turns: []Turn{}}

	game := Game{Engine: &X01Engine{}, Players: &Players{}}

	throw := domain.Throw{Id: 1, Points: 5, Multiplicator: 1, PlayerId: player.Value.Id}
	game.Players.Add(&player)

	game.Engine.RegisterThrow(&throw, game.Players)
	if playerThrows := game.Engine.GetPlayerThrows(&player); len(*playerThrows) == 0 {
		t.Fatalf(`ERROR: Throws for given player should be 1, instead it's zero`)
	}

	if len(game.Players.CurrentPlayer.Turns) != 1 {
		t.Fatalf("ERROR: expected only on turn, since there was till space for adding a throw. instead a new turn was created")
	}*/
}

func TestPlayerLogic(t *testing.T) {
	player := Player{Value: &domain.Player{PlayerName: "1", Id: 1}, Turns: []Turn{}}
	player2 := Player{Value: &domain.Player{PlayerName: "2", Id: 2}, Turns: []Turn{}}
	game := Game{Players: &Players{}, StartingScore: 301, Engine: &X01Engine{}}

	game.Players.Add(&player)
	game.Players.Add(&player2)

	// testing adding players to linked list correctly
	if game.Players.Head.Value.Id != player.Value.Id {
		t.Fatalf(`ERROR: expected id of head should be 1, instead got: %d`, game.Players.Head.Value.Id)
	}
	if game.Players.CurrentPlayer.Value.Id != player.Value.Id {
		t.Fatalf(`ERROR: expected the id of the current player to be 1, instead got: %d`, game.Players.CurrentPlayer.Value.Id)
	}
	if game.Players.Tail.Value.Id != player2.Value.Id {
		t.Fatalf(`ERROR: expected the id of the tail to be 2, instead got: %d`, game.Players.Tail.Value.Id)
	}
	// testing if the bidirectional linking of the list is correct
	if game.Players.Head.Next.Value.Id != player2.Value.Id {
		t.Fatalf(`ERROR: expected the id of the next player after head to be 2, instead got: %d`, game.Players.Head.Next.Value.Id)
	}
	if game.Players.Head.Previous != nil {
		t.Fatalf(`ERROR: expected the previous player from head to be nil, instead got: %v`, game.Players.Head.Previous)
	}
	if game.Players.Tail.Previous.Value.Id != player.Value.Id {
		t.Fatalf(`ERROR: expected the prevous player from the tail to have id 1, instead got: %d`, game.Players.Tail.Previous.Value.Id)
	}
	if game.Players.Tail.Next != nil {
		t.Fatalf(`ERROR: expected the next player from the tail to be nil, instead got: %v`, game.Players.Tail.Next)
	}

	// testing turn switching after three throws
	throw := domain.Throw{Id: 1, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw2 := domain.Throw{Id: 2, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw3 := domain.Throw{Id: 3, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}

	game.Engine.RegisterThrow(&throw, game.Players)
	game.Engine.RegisterThrow(&throw2, game.Players)
	game.Engine.RegisterThrow(&throw3, game.Players)

	if game.Players.CurrentPlayer.Value.Id != player2.Value.Id {
		t.Fatalf(`ERROR: expected the player to change to player2 after three throws. instead got player with id: %d`, game.Players.CurrentPlayer.Value.Id)
	}

	// testing player score calculation after three Throws
	game.Players.CurrentPlayer.Previous.CalculateScore(game.StartingScore)
	if player.Value.Score != 298 {
		t.Fatalf(`ERROR: expected player score after three 1 throws to be 298. instead got: %d`, player.Value.Score)
	}

	game.Engine.RegisterThrow(&throw, game.Players)
	game.Engine.RegisterThrow(&throw2, game.Players)
	game.Engine.RegisterThrow(&throw3, game.Players)

	if game.Players.CurrentPlayer.Value.Id != player.Value.Id {
		t.Fatalf(`ERROR: expected the player to change to player1 after three throws. instead got player with id %d`, game.Players.CurrentPlayer.Value.Id)
	}

	game.Players.CurrentPlayer.Previous.CalculateScore(game.StartingScore)
	if player.Value.Score != 298 {
		t.Fatalf(`ERROR: expected player score after three 1 throws to be 298. instead got: %d`, player.Value.Score)
	}

	var playerTurns []Turn
	for _, turn := range game.Players.CurrentPlayer.Turns {
		playerTurns = append(playerTurns, turn)
	}

	if len(playerTurns) != 1 {
		t.Fatalf(`ERROR: expected the player to have one turn that blongs to him, instead got %d`, len(playerTurns))
	}
}
