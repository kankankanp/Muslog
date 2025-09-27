"use client";

import { useParams } from "next/navigation";
import React, { useEffect, useState, useRef } from "react";
import toast from "react-hot-toast";
import ChatInput from "@/components/community/ChatInput";
import ChatMessage from "@/components/community/ChatMessage";
import { useGetMe } from "@/libs/api/generated/orval/auth/auth";
import { useGetCommunitiesCommunityIdMessages } from "@/libs/api/generated/orval/communities/communities";
import { useWebSocket } from "@/libs/websocket/client";

interface CommunityChatPageProps {
  params: { communityId: string };
}

const CommunityChatPage: React.FC<CommunityChatPageProps> = () => {
  const params = useParams();
  const communityId =
    typeof params.communityId === "string"
      ? params.communityId
      : Array.isArray(params.communityId)
        ? params.communityId[0]
        : "";
  interface Message {
    id: string;
    communityId: string;
    senderId: string;
    content: string;
    createdAt: string;
  }

  const [messages, setMessages] = useState<Message[]>([]);
  const messagesEndRef = useRef<HTMLDivElement>(null);

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
      setMessages(data.messages);
    }
  }, [data]);

  // WebSocket connection
  const wsUrl =
    process.env.NEXT_PUBLIC_WEBSOCKET_URL ||
    (typeof window !== "undefined"
      ? `${window.location.protocol === "https:" ? "wss" : "ws"}://${window.location.hostname}:8080/ws/community/`
      : "wss://localhost:8080/ws/community/");
  const { isConnected, lastMessage, sendMessage } = useWebSocket(
    `${wsUrl}${communityId}`,
    {
      onMessage: (event) => {
        try {
          const newMessage = JSON.parse(event.data);
          setMessages((prevMessages) => [...prevMessages, newMessage]);
        } catch (e) {
          console.error("Failed to parse WebSocket message:", e);
        }
      },
      onError: (event) => {
        console.error("WebSocket error. Please check console.", event);
      },
      onClose: (event) => {
        console.log("Disconnected from chat. Please refresh.", event);
      },
    }
  );

  // Scroll to bottom on new messages
  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
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
      `Error loading chat history: ${error?.message || "Unknown error"}`
    );
    return (
      <div className="text-center text-red-600">
        Error: {error?.message || "Failed to load chat history"}
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
          {isConnected ? "Connected" : "Disconnected"}
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
                msg.id ?? `${msg.senderId ?? "unknown"}-${msg.createdAt ?? idx}`
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
