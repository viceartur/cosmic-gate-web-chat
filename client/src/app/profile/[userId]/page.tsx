"use client";

import { fetchUserById } from "@/actions/users";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

export default function UserProfilePage() {
  const { userId } = useParams();
  const userIdAsString = userId ? String(userId) : "";
  const [userData, setUserData] = useState({
    username: "",
    email: "",
    bio: "",
  });

  useEffect(() => {
    if (!userIdAsString) return;

    const getUserData = async () => {
      const userData: any = await fetchUserById(userIdAsString);
      setUserData(userData);
    };
    getUserData();
  }, [userId]);

  return (
    <section>
      <div className="section-headers">
        <h1>User Profile</h1>
      </div>
      <div className="profile-info">
        <div>
          <label>Username:</label>
          <p>{userData.username}</p>
        </div>
        <div>
          <label>Email:</label>
          <p>{userData.email}</p>
        </div>
        <div>
          <label>Bio:</label>
          <p>{userData.bio}</p>
        </div>
      </div>
    </section>
  );
}
