"use client";

import Messsages from "@/components/messages";
import { useEffect, useState } from "react";

export default function MessagesPage() {
  const [userId, setUserId] = useState<number>(0);
  const [socket, setSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    const handleConnectSocket = () => {
      const userId = Math.round(Math.random() * 1000);
      setUserId(userId);
      const socket = new WebSocket(`ws://localhost:8080/ws/${userId}`);

      socket.onopen = () => {
        const obj = {
          type: "chat-connection",
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
        <h3>Your User ID: {userId}</h3>
      </div>
      <Messsages socket={socket} userId={userId} />
    </>
  );
}
