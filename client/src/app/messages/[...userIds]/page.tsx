"use client";

import Messsages from "@/components/messages";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

export default function MessagesPage() {
  const { userIds }: { userIds: Array<string> } = useParams();
  const [socket, setSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    const handleConnectSocket = () => {
      // Temp solution. Once Auth is created, then user is defined.
      const socket = new WebSocket(`ws://localhost:8080/ws/${userIds[0]}`);

      // Send WebSocket that User connected to the chat.
      socket.onopen = () => {
        const obj = {
          type: "chat-connection",
          recipientId: Number(userIds[1]),
        };
        const jsonString = JSON.stringify(obj);
        socket.send(jsonString);
      };

      socket.onclose = () => {
        console.log("WebSocket connection has been closed.");
      };

      setSocket(socket);
    };
    handleConnectSocket();
  }, []);

  return (
    <>
      <div className="section-headers">
        <h1>Messages Page</h1>
        <h3>Your User ID: {userIds[0]}</h3>
        <h3>Chatting with: {userIds[1]}</h3>
      </div>
      <Messsages socket={socket} userId={userIds[0]} friendId={userIds[1]} />
    </>
  );
}
