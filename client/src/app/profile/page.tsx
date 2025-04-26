"use client";

import { fetchUserById, updateUserData } from "@/actions/users";
import { useSession } from "next-auth/react";
import { FormEvent, useEffect, useState } from "react";

export default function MyProfilePage() {
  const { data: session, update } = useSession();
  const [userData, setUserData] = useState({
    username: "",
    email: "",
    bio: "",
  });
  const [response, setResponse] = useState<string>("");

  useEffect(() => {
    if (!session?.user.id) return;

    const getUserData = async () => {
      const userData: any = await fetchUserById(session?.user.id);
      setUserData(userData);
    };
    getUserData();
  }, [session?.user.id]);

  const submitForm = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    try {
      const formData = new FormData(event.currentTarget);

      const updatedFields: { [key: string]: any } = {};

      const username = formData.get("username") as string;
      const email = formData.get("email") as string;
      const bio = formData.get("bio") as string;
      const password = formData.get("password") as string;
      const confirmPassword = formData.get("confirmPassword") as string;

      // Compare fields and only add changed ones
      if (username && username !== userData.username) {
        updatedFields.username = username;
      }

      if (email && email !== userData.email) {
        updatedFields.email = email;
      }

      if (bio && bio !== userData.bio) {
        updatedFields.bio = bio;
      }

      // Only send password if filled and matches
      if (password) {
        if (password !== confirmPassword) {
          setResponse("Passwords do not match.");
          return;
        }
        updatedFields.password = password;
      }

      // If nothing changed, don't call API
      if (Object.keys(updatedFields).length === 0) {
        setResponse("No changes detected.");
        return;
      }

      const updateResult = await updateUserData(session?.user.id, formData);

      if (updateResult.error) {
        setResponse(updateResult.error);
        return;
      }

      setResponse(updateResult.message);

      // Update the session
      await update({
        user: {
          ...session?.user,
          username: formData.get("username") as string,
          email: formData.get("email") as string,
        },
      });
    } catch (error: any) {
      setResponse(error.message);
    }
  };

  return (
    <section>
      <div className="section-headers">
        <h1>My Profile Page</h1>
      </div>
      <form onSubmit={submitForm} className="profile-form">
        <div>
          <label>Username:</label>
          <input
            type="text"
            name="username"
            placeholder="Username"
            defaultValue={userData.username}
          />
        </div>
        <div>
          <label>Email:</label>
          <input
            type="email"
            name="email"
            placeholder="Email"
            defaultValue={userData.email}
          />
        </div>
        <div>
          <label>Bio:</label>
          <input
            type="text"
            name="bio"
            placeholder="Bio"
            defaultValue={userData.bio}
          />
        </div>
        <div>
          <label>New Password:</label>
          <input type="password" name="password" placeholder="Password" />
        </div>
        <div>
          <label>Confirm New Password:</label>
          <input
            type="password"
            name="confirmPassword"
            placeholder="Password"
          />
        </div>
        {response && <div className="response-message">{response}</div>}
        <button type="submit" className="submit-button">
          Update Profile
        </button>
      </form>
    </section>
  );
}
