package x01

import (
	"fmt"
	"server/internal/triffgonix/dart/engine"
	"server/internal/triffgonix/domain"
	"testing"
)

func TestX01Engine_RegisterThrow(t *testing.T) {
	player := engine.Player{Value: &domain.Player{Id: 1}, Turns: []engine.Turn{}}

	game := engine.Game{Engine: New(301), Players: &engine.Players{}}

	throw := domain.Throw{Id: 1, Points: 5, Multiplicator: 1, PlayerId: player.Value.Id}
	game.Players.Add(&player)

	game.Engine.RegisterThrow(&throw, game.Players)
	if playerThrows := game.Engine.GetPlayerThrows(&player); len(*playerThrows) == 0 {
		t.Fatalf(`ERROR: Throws for given player should be 1, instead it's zero`)
	}

	if len(game.Players.CurrentPlayer.Turns) != 1 {
		t.Fatalf("ERROR: expected only on turn, since there was till space for adding a throw. instead a new turn was created")
	}
}

func TestPlayers_Add(t *testing.T) {
	player := engine.Player{Value: &domain.Player{PlayerName: "1", Id: 1}, Turns: []engine.Turn{}}
	player2 := engine.Player{Value: &domain.Player{PlayerName: "2", Id: 2}, Turns: []engine.Turn{}}
	game := engine.Game{Players: &engine.Players{}, Engine: New(301)}

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
}

func TestX01Engine_RegisterThrow_SwitchPlayer(t *testing.T) {
	player := engine.Player{Value: &domain.Player{PlayerName: "1", Id: 1}, Turns: []engine.Turn{}}
	player2 := engine.Player{Value: &domain.Player{PlayerName: "2", Id: 2}, Turns: []engine.Turn{}}
	game := engine.Game{Players: &engine.Players{}, Engine: New(301)}

	game.Players.Add(&player)
	game.Players.Add(&player2)

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

	game.Engine.RegisterThrow(&throw, game.Players)
	game.Engine.RegisterThrow(&throw2, game.Players)
	game.Engine.RegisterThrow(&throw3, game.Players)

	if game.Players.CurrentPlayer.Value.Id != player.Value.Id {
		t.Fatalf(`ERROR: expected the player to change to player1 after three throws. instead got player with id %d`, game.Players.CurrentPlayer.Value.Id)
	}

	game.Engine.RegisterThrow(&throw, game.Players)
	game.Engine.RegisterThrow(&throw2, game.Players)
	game.Engine.RegisterThrow(&throw3, game.Players)

	if game.Players.CurrentPlayer.Value.Id != player2.Value.Id {
		t.Fatalf(`ERROR: expected the player to change to player2 after three throws. instead got player with id: %d`, game.Players.CurrentPlayer.Value.Id)
	}
}

func TestX01Engine_CalculatePlayerScore(t *testing.T) {
	player := engine.Player{Value: &domain.Player{PlayerName: "1", Id: 1}, Turns: []engine.Turn{}}
	player2 := engine.Player{Value: &domain.Player{PlayerName: "2", Id: 2}, Turns: []engine.Turn{}}
	game := engine.Game{Players: &engine.Players{}, Engine: New(301)}

	game.Players.Add(&player)
	game.Players.Add(&player2)

	throw := domain.Throw{Id: 1, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw2 := domain.Throw{Id: 2, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw3 := domain.Throw{Id: 3, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}

	game.Engine.RegisterThrow(&throw, game.Players)
	game.Engine.RegisterThrow(&throw2, game.Players)
	game.Engine.RegisterThrow(&throw3, game.Players)

	if player.Value.Score != 298 {
		t.Fatalf(`ERROR: expected player score after three 1 throws to be 298. instead got: %d`, game.Players.GetPreviousPlayer().Value.Score)
	}

	throw4 := domain.Throw{Id: 4, Points: 2, Multiplicator: 1, PlayerId: player2.Value.Id}
	throw5 := domain.Throw{Id: 5, Points: 2, Multiplicator: 1, PlayerId: player2.Value.Id}
	throw6 := domain.Throw{Id: 6, Points: 3, Multiplicator: 3, PlayerId: player2.Value.Id}

	game.Engine.RegisterThrow(&throw4, game.Players)
	game.Engine.RegisterThrow(&throw5, game.Players)
	game.Engine.RegisterThrow(&throw6, game.Players)

	if player2.Value.Score != 288 {
		t.Fatalf(`ERROR: expected player score after three 1 throws to be 298. instead got: %d`, game.Players.GetPreviousPlayer().Value.Score)
	}
}

func TestX01Engine_RegisterThrow_CorrectTurnsAmount(t *testing.T) {
	player := engine.Player{Value: &domain.Player{PlayerName: "1", Id: 1}, Turns: []engine.Turn{}}
	player2 := engine.Player{Value: &domain.Player{PlayerName: "2", Id: 2}, Turns: []engine.Turn{}}
	game := engine.Game{Players: &engine.Players{}, Engine: New(301)}

	game.Players.Add(&player)
	game.Players.Add(&player2)

	throw := domain.Throw{Id: 1, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw2 := domain.Throw{Id: 2, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw3 := domain.Throw{Id: 3, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}

	game.Engine.RegisterThrow(&throw, game.Players)
	game.Engine.RegisterThrow(&throw2, game.Players)
	game.Engine.RegisterThrow(&throw3, game.Players)

	var playerTurns []engine.Turn
	for _, turn := range game.Players.GetPreviousPlayer().Turns {
		playerTurns = append(playerTurns, turn)
	}

	if len(playerTurns) != 1 {
		t.Fatalf(`ERROR: expected the player 1 to have one turn that blongs to him, instead got %d`, len(playerTurns))
	}

	game.Engine.RegisterThrow(&throw, game.Players)
	game.Engine.RegisterThrow(&throw2, game.Players)
	game.Engine.RegisterThrow(&throw3, game.Players)

	var player2Turns []engine.Turn
	for _, turn := range game.Players.GetPreviousPlayer().Turns {
		player2Turns = append(player2Turns, turn)
	}

	if len(player2Turns) != 1 {
		t.Fatalf(`ERROR: expected the player 2 to have one turn that blongs to him, instead got %d`, len(player2Turns))
	}

	game.Engine.RegisterThrow(&throw, game.Players)
	game.Engine.RegisterThrow(&throw2, game.Players)
	game.Engine.RegisterThrow(&throw3, game.Players)

	var player1Turns2 []engine.Turn
	for _, turn := range game.Players.GetPreviousPlayer().Turns {
		player1Turns2 = append(player1Turns2, turn)
	}

	if len(player1Turns2) != 2 {
		t.Fatalf(`ERROR: expected the player 1 to have two turn that blongs to him, instead got %d`, len(player1Turns2))
	}
}

func TestX01Engine_HasAnyPlayerWon(t *testing.T) {
	player := engine.Player{Value: &domain.Player{PlayerName: "1", Id: 1, Score: 301}, Turns: []engine.Turn{}}
	player2 := engine.Player{Value: &domain.Player{PlayerName: "2", Id: 2, Score: 0}, Turns: []engine.Turn{}}
	game := engine.Game{Players: &engine.Players{}, Engine: New(301)}

	game.Players.Add(&player)
	game.Players.Add(&player2)

	winningPlayer := game.Engine.HasAnyPlayerWon(game.Players)
	if winningPlayer.Value.Id != player2.Value.Id {
		t.Fatalf(`ERROR: expected player 2 to win. instead player with id: %d won`, winningPlayer.Value.Id)
	}
}

func TestX01Engine_HasAnyPlayerWon_AfterThrows(t *testing.T) {
	player := engine.Player{Value: &domain.Player{PlayerName: "1", Id: 1, Score: 34}, Turns: []engine.Turn{}}
	player2 := engine.Player{Value: &domain.Player{PlayerName: "2", Id: 2, Score: 2}, Turns: []engine.Turn{}}
	game := engine.Game{Players: &engine.Players{}, Engine: New(301)}

	game.Players.Add(&player)
	game.Players.Add(&player2)

	throw := domain.Throw{Id: 1, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw2 := domain.Throw{Id: 2, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw3 := domain.Throw{Id: 3, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}

	game.Engine.RegisterThrow(&throw, game.Players)
	if game.Engine.HasAnyPlayerWon(game.Players) != nil {
		t.Fatal("ERROR: player 1 wasn't supposed to win")
	}
	game.Engine.RegisterThrow(&throw2, game.Players)
	if game.Engine.HasAnyPlayerWon(game.Players) != nil {
		t.Fatal("ERROR: player 1 wasn't supposed to win")
	}
	game.Engine.RegisterThrow(&throw3, game.Players)
	if game.Engine.HasAnyPlayerWon(game.Players) != nil {
		t.Fatal("ERROR: player 1 wasn't supposed to win")
	}

	throw4 := domain.Throw{Id: 4, Points: 255, Multiplicator: 1, PlayerId: player2.Value.Id}
	throw5 := domain.Throw{Id: 5, Points: 37, Multiplicator: 1, PlayerId: player2.Value.Id}
	throw6 := domain.Throw{Id: 6, Points: 3, Multiplicator: 3, PlayerId: player2.Value.Id}

	game.Engine.RegisterThrow(&throw4, game.Players)
	if game.Engine.HasAnyPlayerWon(game.Players) != nil {
		t.Fatal("ERROR: player 2 was not supposed to win yet")
	}
	game.Engine.RegisterThrow(&throw5, game.Players)
	if game.Engine.HasAnyPlayerWon(game.Players) != nil {
		t.Fatal("ERROR: player 2 was not supposed to win yet")
	}
	game.Engine.RegisterThrow(&throw6, game.Players)
	if game.Engine.HasAnyPlayerWon(game.Players) == nil {
		t.Fatal("ERROR: player 2 was supposed to win")
	}
}

func TestX01Engine_UndoThrow(t *testing.T) {
	player := engine.Player{Value: &domain.Player{PlayerName: "1", Id: 1, Score: 34}, Turns: []engine.Turn{}}
	game := engine.Game{Players: &engine.Players{}, Engine: New(301)}

	game.Players.Add(&player)

	throw := domain.Throw{Id: 1, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}

	game.Engine.RegisterThrow(&throw, game.Players)
	playerThrows := game.Engine.GetPlayerThrows(&player)
	if len(*playerThrows) != 1 {
		t.Fatal("ERROR: expected player throws to contain one throw")
	}

	game.Engine.UndoLastThrow(game.Players)
	playerThrows2 := game.Engine.GetPlayerThrows(&player)

	fmt.Printf("trace: playerThrows2 - %+v\n", playerThrows2)
	if len(*playerThrows2) != 0 {
		t.Fatal("ERROR: expected player throws to contain no throws after undo")
	}
}

func TestX01Engine_UndoThrow_WhenSecondPlayerTurn(t *testing.T) {
	player := engine.Player{Value: &domain.Player{PlayerName: "1", Id: 1, Score: 34}, Turns: []engine.Turn{}}
	player2 := engine.Player{Value: &domain.Player{PlayerName: "2", Id: 2, Score: 2}, Turns: []engine.Turn{}}
	game := engine.Game{Players: &engine.Players{}, Engine: New(301)}

	game.Players.Add(&player)
	game.Players.Add(&player2)

	throw := domain.Throw{Id: 1, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw2 := domain.Throw{Id: 2, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw3 := domain.Throw{Id: 3, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}

	game.Engine.RegisterThrow(&throw, game.Players)
	game.Engine.RegisterThrow(&throw2, game.Players)
	game.Engine.RegisterThrow(&throw3, game.Players)

	if len(*game.Engine.GetPlayerThrows(&player)) != 3 {
		t.Fatal("ERROR: expected the first player to have three throws registred")
	}

	game.Engine.RegisterThrow(&throw, game.Players)

	if len(*game.Engine.GetPlayerThrows(&player2)) != 1 {
		t.Fatal("ERROR: expected the second player to have one throw registred")
	}

	game.Engine.UndoLastThrow(game.Players)

	if len(*game.Engine.GetPlayerThrows(&player2)) != 0 {
		t.Fatal("ERROR: expected the second player to have zero throws registred after undo")
	}

	game.Engine.UndoLastThrow(game.Players)

	if len(*game.Engine.GetPlayerThrows(&player)) != 2 {
		t.Fatal("ERROR: expected the first player to have two throws registred after undo")
	}

	if game.Players.CurrentPlayer.Value.Id != player.Value.Id {
		t.Fatal("ERROR: expected the first player to be current player after undo")
	}
}

func TestX01Engine_OverThrow_Scenario(t *testing.T) {
	player := engine.Player{Value: &domain.Player{PlayerName: "1", Id: 1, Score: 0}, Turns: []engine.Turn{}}
	player2 := engine.Player{Value: &domain.Player{PlayerName: "2", Id: 2, Score: 0}, Turns: []engine.Turn{}}
	game := engine.Game{Players: &engine.Players{}, Engine: New(301)}

	game.Players.Add(&player)
	game.Players.Add(&player2)

	throw := domain.Throw{Id: 1, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw2 := domain.Throw{Id: 2, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw3 := domain.Throw{Id: 3, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw4 := domain.Throw{Id: 4, Points: 302, Multiplicator: 1, PlayerId: player.Value.Id}
	game.Engine.RegisterThrow(&throw, game.Players)
	game.Engine.RegisterThrow(&throw2, game.Players)
	game.Engine.RegisterThrow(&throw3, game.Players)

	if game.Players.CurrentPlayer.Value.Id != player2.Value.Id {
		t.Fatalf("ERROR: expected player two to be the current player")
	}
	game.Engine.RegisterThrow(&throw4, game.Players)
	if player2.Turns[0].Sum() != 0 {
		t.Fatalf("ERROR: expected player two on a overthrow to have a zero points in turn. instead got: %d",
			player2.Turns[0].Sum())
	}
	if game.Players.CurrentPlayer.Value.Id != player.Value.Id {
		t.Fatalf("ERROR: expected player one to be current player after overthrow")
	}

	throw5 := domain.Throw{Id: 1, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw6 := domain.Throw{Id: 2, Points: 302, Multiplicator: 1, PlayerId: player.Value.Id}

	game.Engine.RegisterThrow(&throw5, game.Players)
	game.Engine.RegisterThrow(&throw6, game.Players)
	if player.Turns[1].Sum() != 1 {
		t.Fatalf("ERROR: expected player one on a overthrow to have one point in turn. instead got: %d", player.Turns[1].Sum())
	}
	if game.Players.CurrentPlayer.Value.Id != player2.Value.Id {
		t.Fatalf("ERROR: expected player two to be the current player")
	}
}

func TestX01Engine_OverThrow_NoSpaceInTurnScenario(t *testing.T) {
	player := engine.Player{Value: &domain.Player{PlayerName: "1", Id: 1, Score: 0}, Turns: []engine.Turn{}}
	player2 := engine.Player{Value: &domain.Player{PlayerName: "2", Id: 2, Score: 0}, Turns: []engine.Turn{}}
	game := engine.Game{Players: &engine.Players{}, Engine: New(301)}

	game.Players.Add(&player)
	game.Players.Add(&player2)

	throw := domain.Throw{Id: 1, Points: 298, Multiplicator: 1, PlayerId: player.Value.Id}
	throw2 := domain.Throw{Id: 2, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}
	throw3 := domain.Throw{Id: 3, Points: 1, Multiplicator: 1, PlayerId: player.Value.Id}

	game.Engine.RegisterThrow(&throw, game.Players)
	game.Engine.RegisterThrow(&throw2, game.Players)
	game.Engine.RegisterThrow(&throw3, game.Players)

	throw4 := domain.Throw{Id: 4, Points: 1, Multiplicator: 1, PlayerId: player2.Value.Id}
	throw5 := domain.Throw{Id: 3, Points: 1, Multiplicator: 1, PlayerId: player2.Value.Id}
	throw6 := domain.Throw{Id: 3, Points: 1, Multiplicator: 1, PlayerId: player2.Value.Id}

	game.Engine.RegisterThrow(&throw4, game.Players)
	game.Engine.RegisterThrow(&throw5, game.Players)
	game.Engine.RegisterThrow(&throw6, game.Players)

	throw7 := domain.Throw{Id: 3, Points: 4, Multiplicator: 1, PlayerId: player.Value.Id}
	game.Engine.RegisterThrow(&throw7, game.Players)

	if game.Players.CurrentPlayer.Value.Id != player2.Value.Id {
		t.Fatalf("ERROR: expected player one to be the current player")
	}

	if player.Turns[1].Sum() != 0 {
		t.Fatalf("ERROR: expected player one on a overthrow to have a zero points in turn. instead got: %d",
			player.Turns[1].Sum())
	}
}
