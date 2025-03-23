export interface NewGamePlayer {
  id: number;
  name: string;
}

export interface GameStatePlayer {
  id: number;
  name: string;
  averagePoints: number;
  score: number;
}

export interface GameStatePlayers {
  allPlayers: GameStatePlayer[];
  currentPlayer: GameStatePlayer;
}

export interface GameStateContent {
  id: number;
  name: string;
  gameMode: string;
  players: GameStatePlayers;
}

export interface NewGameContent {
  id: number;
  name: string;
  gameMode: string;
  players: NewGamePlayer[];
}

export interface IncomingSocketMessage {
  type: string;
  content: GameStateContent[] | NewGameContent;
}
