
export class Player {
  id: number
  name: string

  constructor(id: number, name: string) {
    this.id = id;
    this.name = name;
  }
}

export class Game {
  name: string
  gameMode: string
  startingScore: number
  players: Player[]

  constructor(name: string, gameMode: string, startingScore: number, players: Player[]) {
    this.name = name;
    this.gameMode = gameMode;
    this.startingScore = startingScore;
    this.players = players;

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
  type: string
  content: ThrowContent | HandshakeContent | UndoThrowContent

  constructor(type: string, content: ThrowContent | HandshakeContent | UndoThrowContent) {
    this.type = type;
    this.content = content;
  }
}

