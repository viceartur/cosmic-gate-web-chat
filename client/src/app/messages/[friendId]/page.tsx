"use client";

import Messages from "@/components/messages";
import { useSession } from "next-auth/react";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

export default function MessagesPage() {
  const { data: session } = useSession();
  const { friendId }: { friendId: string } = useParams();
  const [socket, setSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    if (!session?.user.id) return;

    const socket = new WebSocket(`ws://localhost:8080/ws/${session.user.id}`);

    // Send WebSocket that User connected to the chat.
    socket.onopen = () => {
      const obj = {
        type: "chat-connection",
        recipientId: friendId,
      };
      const jsonString = JSON.stringify(obj);
      socket.send(jsonString);
    };

    socket.onclose = () => {
      console.log("WebSocket connection has been closed.");
    };

    setSocket(socket);

    // Close WebSocket once the component is unmounted
    return () => {
      socket.close();
    };
  }, [session?.user.id, friendId]);

  return (
    <>
      <div className="section-headers">
        <h1>Messages Page</h1>
        <h3>Hey {session?.user.username}</h3>
        <h3>You are chatting with: {friendId}</h3>
      </div>
      <Messages socket={socket} userId={session?.user.id} friendId={friendId} />
    </>
  );
}
