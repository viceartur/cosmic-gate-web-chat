"use client";

import Link from "next/link";
import { signOut, useSession } from "next-auth/react";
import { useEffect, useState } from "react";
import { fetchUserById } from "@/actions/users";

export function NavBar() {
  const { data: session } = useSession();
  const [numFriendRequests, setNumFriendRequests] = useState(0);

  useEffect(() => {
    if (!session?.user?.id) return;

    // Create WebSocket instance:
    // - The instance available within the NavBar Component Scope, so it's closed in other components
    // - Doesn't work in Users and Messages pages since they have own WebSocket instances
    const ws = new WebSocket(`ws://localhost:8080/ws/${session.user.id}`);

    // Get a Friend Request from a User
    ws.onmessage = (event: MessageEvent) => {
      const socketMsg = JSON.parse(event.data);
      console.log("socketMsg", socketMsg);
      if (socketMsg.type === "friend-requests") {
        setNumFriendRequests(Number(socketMsg.data));
      }
    };

    ws.onclose = () => {
      console.log("WebSocket connection has been closed.");
    };

    // Check the Number of Friend Requests if WebSocket wasn't sent
    const getNumberFriendRequests = async () => {
      const user: any = await fetchUserById(session.user.id);
      if (user.friendRequests.length) {
        setNumFriendRequests(user.friendRequests.length);
      }
    };
    getNumberFriendRequests();

    // Close WebSocket once the component is unmounted
    return () => {
      ws.close();
    };
  }, [session?.user?.id]);

  if (!session?.user) {
    return null;
  }

  return (
    <nav>
      <Link href="/">Main</Link>
      <Link href="/profile">Profile</Link>
      <Link href="/users">People</Link>
      <Link href="/friend-requests">
        Friend Requests {numFriendRequests ? ` (${numFriendRequests})` : ""}
      </Link>
      <Link href="/friends">Friends</Link>
      {session?.user && <button onClick={() => signOut()}>Sign Out</button>}
    </nav>
  );
}
