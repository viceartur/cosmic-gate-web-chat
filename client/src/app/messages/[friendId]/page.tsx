"use client";

import Messages from "@/components/messages";
import { useSession } from "next-auth/react";
import { useParams, useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";

export default function MessagesPage() {
  const friendName = useSearchParams().get("username");
  const { friendId }: { friendId: string } = useParams();
  const { data: session } = useSession();
  const [socket, setSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    if (!session?.user.id) return;

    // Create WebSocket Instance
    const ws = new WebSocket(`ws://localhost:8080/ws/${session.user.id}`);

    // Send WebSocket to a Friend that the User connected to the chat
    ws.onopen = () => {
      const obj = {
        type: "chat-connection",
        recipientId: friendId,
      };
      const jsonString = JSON.stringify(obj);
      ws.send(jsonString);
    };

    ws.onclose = () => {
      console.log("WebSocket connection has been closed.");
    };

    setSocket(ws);

    // Close WebSocket once the component is unmounted
    return () => {
      ws.close();
    };
  }, [session?.user.id]);

  return (
    <section>
      <div className="section-headers">
        <h1>Chat with {friendName}</h1>
      </div>
      <Messages socket={socket} userId={session?.user.id} friendId={friendId} />
    </section>
  );
}
