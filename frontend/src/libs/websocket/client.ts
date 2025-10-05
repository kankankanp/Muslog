import { useEffect, useRef, useState, useCallback } from "react";

interface WebSocketOptions {
  onOpen?: (event: Event) => void;
  onMessage?: (event: MessageEvent) => void;
  onClose?: (event: CloseEvent) => void;
  onError?: (event: Event) => void;
}

export const useWebSocket = (url: string, options?: WebSocketOptions) => {
  const [isConnected, setIsConnected] = useState(false);
  const [lastMessage, setLastMessage] = useState<MessageEvent | null>(null);
  const wsRef = useRef<WebSocket | null>(null);

  const sendMessage = useCallback((message: string) => {
    if (wsRef.current && wsRef.current.readyState === WebSocket.OPEN) {
      wsRef.current.send(message);
    } else {
      console.warn("WebSocket is not open. Message not sent:", message);
    }
  }, []);

  useEffect(() => {
    if (!url) return;

    const ws = new WebSocket(url);
    wsRef.current = ws;

    ws.onopen = (event) => {
      setIsConnected(true);
      options?.onOpen?.(event);
      console.log("WebSocket opened:", url);
    };

    ws.onmessage = (event) => {
      setLastMessage(event);
      options?.onMessage?.(event);
    };

    ws.onclose = (event) => {
      setIsConnected(false);
      options?.onClose?.(event);
      console.log("WebSocket closed:", url, event);
    };

    ws.onerror = (event) => {
      options?.onError?.(event);
      console.error("WebSocket error:", url, event);
    };

    return () => {
      if (wsRef.current) {
        wsRef.current.close();
      }
    };
  }, [options, url]);

  return { ws: wsRef.current, isConnected, lastMessage, sendMessage };
};
