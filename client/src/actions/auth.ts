"use server";

import { API } from "@/utils/constants";

// Authenticate an existing User
// Check that email and password provided are correct
export const authUser = async (email: string, password: string) => {
  const response = await fetch(`${API}/auth`, {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });
  if (!response.ok) {
    return null;
  }

  const user = await response.json();
  const userData = {
    id: user.id,
    email: user.email,
    username: user.username,
  };
  return userData;
};

// Sign Up a new User
export const signUp = async ({
  email,
  password,
  username,
}: {
  email: string;
  password: string;
  username: string;
}) => {
  // Verify that an Account doesn't exist
  const foundUser = await fetch(`${API}/users?email=${email}`);
  if (foundUser.ok) {
    throw new Error("User already exists. Please Log In using credentials");
  }

  // Create an Account
  const newUser = await fetch(`${API}/users`, {
    method: "POST",
    body: JSON.stringify({ email, username, password }),
  });
  if (!newUser.ok) {
    throw new Error("Unable to create an Account");
  }

  const newUserData: any = newUser.json();
  const newUserId = newUserData.InsertedID;

  return { userId: newUserId };
};
