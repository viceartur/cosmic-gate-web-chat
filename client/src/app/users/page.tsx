"use client";

import { useSession } from "next-auth/react";
import { redirect } from "next/navigation";
import { useEffect, useState } from "react";

import { fetchUserFriends, fetchUsers } from "@/actions/users";

export default function FriendsPage() {
  const { data: session } = useSession();
  const [userFriends, setUserFriends] = useState([]);
  const [users, setUsers] = useState([]);

  useEffect(() => {
    const getFriends = async () => {
      if (!session?.user.id) return;
      const userFriends = await fetchUserFriends(session.user.id);
      const allUsers = await fetchUsers(session.user.id);
      setUserFriends(userFriends);
      setUsers(allUsers);
    };
    getFriends();
  }, [session?.user.id]);

  const sendFriendRequest = async (newFriendId: string) => {};

  return (
    <section>
      <div className="section-headers">
        <h1>Chat Users</h1>
      </div>
      <div className="users-list">
        {users.map((user: any, i) => (
          <div key={i} className="user-card">
            <span className="user-name">{user.username}</span>
            <div className="user-actions__buttons">
              <button onClick={() => redirect(`/profile/${user.id}`)}>
                Profile
              </button>
              {userFriends.some((friend: any) => friend.id === user.id) ? (
                <span className="friend-label">Your Friend</span>
              ) : (
                <button onClick={() => sendFriendRequest(user.id)}>
                  Add Friend
                </button>
              )}
            </div>
          </div>
        ))}
      </div>
    </section>
  );
}
