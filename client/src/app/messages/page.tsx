"use client";

import Messsages from "@/components/messages";
import { useEffect, useState } from "react";

export default function MessagesPage() {
  const [userId, setUserId] = useState<number>(1);
  const [socket, setSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    const handleConnectSocket = () => {
      const socket = new WebSocket(`ws://localhost:8080/ws/${userId}`);

      socket.onopen = () => {
        const obj = {
          type: "chat-connection",
          message: "user connected",
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
      <h1>Messages Page</h1>
      <Messsages socket={socket} />
    </>
  );
}
