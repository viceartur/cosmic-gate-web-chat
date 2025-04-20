"use client";

import {
  acceptFriendRequest,
  declineFriendRequest,
  fetchUserFriendRequests,
} from "@/actions/users";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";

export default function FriendRequestsPage() {
  const { data: session } = useSession();
  const [friendRequestSenders, setFriendRequestSenders] = useState([]);
  const [handledRequests, setHandledRequests] = useState<
    Record<string, "accepted" | "declined">
  >({});
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getFriendRequests = async () => {
      if (!session?.user.id) return;

      const requests = await fetchUserFriendRequests(session.user.id);
      setFriendRequestSenders(requests);
      setLoading(false);
    };
    getFriendRequests();
  }, [session?.user.id]);

  const handleFriendAccept = async (friendId: string) => {
    if (!session?.user.id) return;
    await acceptFriendRequest(session.user.id, friendId);
    setHandledRequests((prev) => ({ ...prev, [friendId]: "accepted" }));
  };

  const handleFriendDecline = async (friendId: string) => {
    if (!session?.user.id) return;
    await declineFriendRequest(session.user.id, friendId);
    setHandledRequests((prev) => ({ ...prev, [friendId]: "declined" }));
  };

  return (
    <section>
      <div className="section-headers">
        <h1>Your Friend Requests</h1>
      </div>
      {loading ? (
        <p>Loading friend requests...</p>
      ) : friendRequestSenders.length === 0 ? (
        <p>No new friend requests</p>
      ) : (
        <div className="users-list">
          {friendRequestSenders.map((sender: any) => {
            const status = handledRequests[sender.id];
            return (
              <div key={sender.id} className="user-card">
                <span className="user-name">{sender.username}</span>
                <div className="user-actions__buttons">
                  {status === "accepted" ? (
                    <span className="friend-label">User added</span>
                  ) : status === "declined" ? (
                    <span className="friend-label">User declined</span>
                  ) : (
                    <>
                      <button onClick={() => handleFriendAccept(sender.id)}>
                        Accept
                      </button>
                      <button onClick={() => handleFriendDecline(sender.id)}>
                        Decline
                      </button>
                    </>
                  )}
                </div>
              </div>
            );
          })}
        </div>
      )}
    </section>
  );
}
