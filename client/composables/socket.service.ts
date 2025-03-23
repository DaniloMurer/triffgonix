export const useSocketService = () => {
  const listenOnMessage = (handler: (event: MessageEvent) => void) => {
    const webSocket = new WebSocket('ws://localhost:8080/ws/dart');
    webSocket.onmessage = handler;
  };

  return {
    listenOnMessage,
  };
};
