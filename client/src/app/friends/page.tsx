"use client";

import { fetchUserFriends } from "@/actions/users";
import { useSession } from "next-auth/react";
import { redirect } from "next/navigation";
import { useEffect, useState } from "react";

export default function FriendsPage() {
  const { data: session } = useSession();
  const [friends, setFriends] = useState([]);

  useEffect(() => {
    const getFriends = async () => {
      if (!session?.user.id) return;
      const userFriends = await fetchUserFriends(session.user.id);
      setFriends(userFriends);
    };
    getFriends();
  }, [session?.user.id]);

  return (
    <section>
      <div className="section-headers">
        <h1>Your Friends List</h1>
      </div>
      <div className="friends-list">
        {friends.map((friend: any, i) => (
          <div key={i} className="friend-card">
            <span className="friend-name">{friend.username}</span>
            <button
              className="chat-button"
              onClick={() =>
                redirect(`/messages/${friend.id}?username=${friend.username}`)
              }
            >
              Chat
            </button>
          </div>
        ))}
      </div>
    </section>
  );
}
