import React from "react";
import { components } from "@/libs/api/generated/openapi-types";

type Message = components["schemas"]["Message"];

interface ChatMessageProps {
  message: Message;
  isOwnMessage: boolean;
}

const ChatMessage: React.FC<ChatMessageProps> = ({ message, isOwnMessage }) => {
  const messageClass = isOwnMessage
    ? "bg-blue-500 text-white self-end rounded-br-none"
    : "bg-gray-200 text-gray-800 self-start rounded-bl-none dark:bg-gray-700 dark:text-gray-200";

  return (
    <div className={`flex flex-col p-3 rounded-lg max-w-[70%] ${messageClass}`}>
      <div className="text-xs font-semibold mb-1">
        {isOwnMessage ? "You" : message.senderId}
      </div>
      <div className="text-sm break-words">{message.content}</div>
      <div className="text-xs mt-1 opacity-75">
        {new Date(message.createdAt).toLocaleTimeString([], {
          hour: "2-digit",
          minute: "2-digit",
        })}
      </div>
    </div>
  );
};

export default ChatMessage;
