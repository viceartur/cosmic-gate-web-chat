"use client";

import { useSession } from "next-auth/react";
import { redirect } from "next/navigation";
import { useEffect, useState } from "react";

import {
  fetchUserFriends,
  fetchUsers,
  sendFriendRequest,
} from "@/actions/users";

export default function UsersPage() {
  const { data: session } = useSession();
  const [userFriends, setUserFriends] = useState<{ id: string }[]>([]);
  const [users, setUsers] = useState<
    {
      id: string;
      username: string;
      friendRequests: string[];
      friends: string[];
    }[]
  >([]);
  const [socket, setSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    if (!session?.user.id) return;

    // Open WebSocket
    const ws = new WebSocket(`ws://localhost:8080/ws/${session.user.id}`);
    ws.onclose = () => {
      console.log("WebSocket connection has been closed.");
    };
    setSocket(ws);

    // Fetch User's Friends and all Users
    const getUsers = async () => {
      const userFriends = await fetchUserFriends(session.user.id);
      const allUsers = await fetchUsers(session.user.id);
      setUserFriends(userFriends);
      setUsers(allUsers);
    };
    getUsers();

    // Close WebSocket once the component is unmounted
    return () => {
      ws.close();
    };
  }, [session?.user.id]);

  // Send a Friend Request to a User
  const handleSendFriendRequest = async (newFriendId: string) => {
    const userId = session?.user.id;
    if (!userId) return console.error("User session not defined");

    try {
      await sendFriendRequest(userId, newFriendId);

      // Send Socket to a User
      if (socket) {
        const wsObj = {
          type: "friend-request-sent",
          recipientId: newFriendId,
        };
        const wsJSON = JSON.stringify(wsObj);
        socket.send(wsJSON);
      }

      setUsers((prevUsers) =>
        prevUsers.map((user) =>
          user.id === newFriendId
            ? {
                ...user,
                friendRequests: [...(user.friendRequests || []), userId],
              }
            : user
        )
      );
    } catch (error) {
      console.error("Failed to send friend request:", error);
    }
  };

  return (
    <section>
      <div className="section-headers">
        <h1>Chat Users</h1>
      </div>
      <div className="users-list">
        {users.map((user, i) => (
          <div key={i} className="user-card">
            <span className="user-name">{user.username}</span>
            <div className="user-actions__buttons">
              <button onClick={() => redirect(`/profile/${user.id}`)}>
                Profile
              </button>
              {
                // Check if a User has a friend request from the Auth User
                user.friendRequests.some(
                  (friendReq: any) => friendReq === session?.user.id
                ) ? (
                  <span className="friend-label">Request Sent</span>
                ) : // Check if a User already in Auth User's Friends
                userFriends.some((friend: any) => friend.id === user.id) ? (
                  <span className="friend-label">Friends</span>
                ) : (
                  <button onClick={() => handleSendFriendRequest(user.id)}>
                    Add Friend
                  </button>
                )
              }
            </div>
          </div>
        ))}
      </div>
    </section>
  );
}
