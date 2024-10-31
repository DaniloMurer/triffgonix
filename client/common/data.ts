
export class PlayerDto {
  id: number;
  name: string;

  constructor(id: number, name: string) {
    this.id = id;
    this.name = name;
  }
}

export class GameDto {
  name: string
  gameMode: string
  startingScore: number
  players: PlayerDto[]

  constructor(name: string, gameMode: string, startingScore: number, players: PlayerDto[]) {
    this.name = name;
    this.gameMode = gameMode;
    this.startingScore = startingScore;
    this.players = players;

  }
}

/**
 * Type for player sent by websocket
 */
export class PlayerState {
  id: number
  playerName: string
  score: number
  averagePoints: number

  constructor(id: number, playerName: string, score: number, averagePoints: number) {
    this.id = id;
    this.playerName = playerName;
    this.score = score;
    this.averagePoints = averagePoints;
  }
}

/**
 * Type for game sent by websocket
 */
export class GameState {
  allPlayers: PlayerState[]
  currentPlayer: PlayerState

  constructor(allPlayers: PlayerState[], currentPlayer: PlayerState) {
    this.allPlayers = allPlayers;
    this.currentPlayer = currentPlayer;
  }
}

export class ThrowContent {
  points: number
  multiplicator: number

  constructor(points: number, multiplicator: number) {
    this.points = points;
    this.multiplicator = multiplicator;
  }
}

export class HandshakeContent { }

export class UndoThrowContent { }

export class SocketMessage {
  type: MessageType
  content: ThrowContent | HandshakeContent | UndoThrowContent | GameDto

  constructor(type: MessageType, content: ThrowContent | HandshakeContent | UndoThrowContent | GameDto) {
    this.type = type;
    this.content = content;
  }
}

export enum MessageType {
  THROW = "throw",
  HANDSHAKE = "handshake",
  SAVE = "save",
  UNDO_THROW = "undo-throw",
  NEW_GAME = "new-game",
  GAME_STATE = "game-state"
}
