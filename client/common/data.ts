export type Player = {
  id: number,
  name: string
}

export type Game = {
  name: string,
  gameMode: string,
  startingScore: number,
  players: Player[]
}

export type PlayerState = {
  id: number,
  playerName: string,
  score: number,
  averagePoints: number
}

export type GameState = {
  allPlayers: PlayerState[],
  currentPlayer: PlayerState
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

// export class SocketMessage {
//   type: MessageType
//   content: ThrowContent | HandshakeContent | UndoThrowContent | GameDto
//
//   constructor(type: MessageType, content: ThrowContent | HandshakeContent | UndoThrowContent | GameDto) {
//     this.type = type;
//     this.content = content;
//   }
// }

export enum MessageType {
  THROW = "throw",
  HANDSHAKE = "handshake",
  SAVE = "save",
  UNDO_THROW = "undo-throw",
  NEW_GAME = "new-game",
  GAME_STATE = "game-state"
}
