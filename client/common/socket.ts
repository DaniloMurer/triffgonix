
/**
 * @param hub the game id to connect to
 * @return {WebSocket} the socket instance
 */
export function connectToSocket(hub: string): WebSocket {
  const socket = new WebSocket(`http://localhost:8080/ws/dart/${hub}`);
  return socket;
}


