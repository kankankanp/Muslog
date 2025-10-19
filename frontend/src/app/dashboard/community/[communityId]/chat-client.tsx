'use client';

import { useParams } from 'next/navigation';
import React, {
  useEffect,
  useState,
  useRef,
  useMemo,
  useCallback,
} from 'react';
import toast from 'react-hot-toast';
import ChatInput from '@/components/community/ChatInput';
import ChatMessage from '@/components/community/ChatMessage';
import { useGetMe } from '@/libs/api/generated/orval/auth/auth';
import { useGetCommunitiesCommunityIdMessages } from '@/libs/api/generated/orval/communities/communities';
import { useWebSocket } from '@/libs/websocket/client';

interface CommunityChatPageProps {
  params: { communityId: string };
}

const CommunityChatPage: React.FC<CommunityChatPageProps> = () => {
  const params = useParams();
  const communityId =
    typeof params.communityId === 'string'
      ? params.communityId
      : Array.isArray(params.communityId)
        ? params.communityId[0]
        : '';
  interface Message {
    id: string;
    communityId: string;
    senderId: string;
    content: string;
    createdAt: string;
  }

  const [messages, setMessages] = useState<Message[]>([]);
  const messagesEndRef = useRef<HTMLDivElement>(null);
  const messageIdsRef = useRef<Set<string>>(new Set());

  const { data: me } = useGetMe();

  // Fetch historical messages
  const { data, isLoading, isError, error } =
    useGetCommunitiesCommunityIdMessages(communityId) as {
      data?: { messages: Message[] };
      isLoading: boolean;
      isError: boolean;
      error?: { message?: string };
    };

  useEffect(() => {
    if (data?.messages) {
      const ids = new Set<string>();
      data.messages.forEach((msg) => {
        if (msg.id) {
          ids.add(msg.id);
        }
      });
      messageIdsRef.current = ids;
      setMessages(data.messages);
    }
  }, [data]);

  const parseWebSocketPayload = useCallback((payload: string): Message[] => {
    const trimmed = payload.trim();
    if (!trimmed) return [];

    const results: Message[] = [];
    let buffer = '';
    let depth = 0;
    let inString = false;
    let escape = false;

    for (let i = 0; i < trimmed.length; i++) {
      const char = trimmed[i];
      buffer += char;

      if (inString) {
        if (escape) {
          escape = false;
        } else if (char === '\\') {
          escape = true;
        } else if (char === '"') {
          inString = false;
        }
        continue;
      }

      if (char === '"') {
        inString = true;
      } else if (char === '{') {
        depth++;
      } else if (char === '}') {
        depth--;
      }

      if (depth == 0) {
        const candidate = buffer.trim();
        if (candidate) {
          try {
            results.push(JSON.parse(candidate));
          } catch (parseError) {
            console.error('Failed to parse WebSocket message chunk:', {
              candidate,
              parseError,
            });
          }
        }
        buffer = '';
      }
    }

    return results;
  }, []);

  // WebSocket connection
  const wsUrl =
    process.env.NEXT_PUBLIC_WEBSOCKET_URL ||
    (typeof window !== 'undefined'
      ? `${window.location.protocol === 'https:' ? 'wss' : 'ws'}://${window.location.hostname}:8080/ws/community/`
      : 'wss://localhost:8080/ws/community/');
  const websocketHandlers = useMemo(
    () => ({
      onMessage: (event: MessageEvent) => {
        const handlePayload = (payload: string) => {
          const parsed = parseWebSocketPayload(payload);
          if (parsed.length > 0) {
            const deduped: Message[] = [];
            parsed.forEach((msg) => {
              const key =
                msg.id || `${msg.senderId ?? 'unknown'}-${msg.createdAt ?? ''}`;
              if (!key || messageIdsRef.current.has(key)) {
                return;
              }
              messageIdsRef.current.add(key);
              deduped.push({ ...msg, id: msg.id || key });
            });
            if (deduped.length > 0) {
              setMessages((prevMessages) => [...prevMessages, ...deduped]);
            }
          }
        };

        if (typeof event.data === 'string') {
          handlePayload(event.data);
        } else if (event.data instanceof Blob) {
          event.data
            .text()
            .then(handlePayload)
            .catch((blobError) =>
              console.error('Failed to read Blob WebSocket message:', blobError)
            );
        } else {
          console.warn('Unsupported WebSocket payload type', event.data);
        }
      },
      onError: (event: Event) => {
        console.error('WebSocket error. Please check console.', event);
      },
      onClose: (event: CloseEvent) => {
        console.log('Disconnected from chat. Please refresh.', event);
      },
    }),
    [parseWebSocketPayload]
  );

  const { isConnected, lastMessage, sendMessage } = useWebSocket(
    `${wsUrl}${communityId}`,
    websocketHandlers
  );

  // Scroll to bottom on new messages
  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [messages]);

  const handleSendMessage = (content: string) => {
    // In a real app, senderId would come from authenticated user context
    // For now, backend assigns a guest ID.
    sendMessage(content);
  };

  if (isLoading) {
    return (
      <div className="text-center text-gray-600 dark:text-gray-300">
        Loading chat history...
      </div>
    );
  }

  if (isError) {
    toast.error(
      `Error loading chat history: ${error?.message || 'Unknown error'}`
    );
    return (
      <div className="text-center text-red-600">
        Error: {error?.message || 'Failed to load chat history'}
      </div>
    );
  }

  return (
    <div className="flex flex-col h-screen bg-gray-50 dark:bg-gray-900">
      <div className="flex-none p-4 border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 shadow-sm">
        <h1 className="text-2xl font-bold text-gray-900 dark:text-white">
          Community: {communityId}
        </h1>
        <p className="text-sm text-gray-600 dark:text-gray-400">
          {isConnected ? 'Connected' : 'Disconnected'}
        </p>
      </div>
      <div className="flex-grow overflow-y-auto p-4 space-y-4">
        {messages.length === 0 ? (
          <div className="text-center text-gray-500 dark:text-gray-400">
            No messages yet. Start the conversation!
          </div>
        ) : (
          messages.map((msg, idx) => (
            // TODO: Replace with actual user ID from context for isOwnMessage check
            <ChatMessage
              key={
                msg.id ?? `${msg.senderId ?? 'unknown'}-${msg.createdAt ?? idx}`
              }
              message={msg}
              isOwnMessage={Boolean(me?.id) && msg.senderId === me!.id}
            />
          ))
        )}
        <div ref={messagesEndRef} />
      </div>

      <div className="flex-none p-4 bg-white dark:bg-gray-800 border-t border-gray-200 dark:border-gray-700">
        <ChatInput onSendMessage={handleSendMessage} disabled={!isConnected} />
      </div>
    </div>
  );
};

export default CommunityChatPage;
