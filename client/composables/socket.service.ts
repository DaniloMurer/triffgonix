import { WEBSOCKET_BASE_PATH } from '~/utils/app.config';

let webSocket: WebSocket;

export const useSocketService = () => {
  const connectToSocket = (path: string) => {
    webSocket = new WebSocket(new URL(path, WEBSOCKET_BASE_PATH).toString());
  };

  const listenOnMessage = (handler: (event: MessageEvent) => void) => {
    webSocket.addEventListener('message', handler);
  };

  return {
    connectToSocket,
    listenOnMessage,
  };
};
