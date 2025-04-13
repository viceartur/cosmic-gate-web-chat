import { API } from "@/utils/constants";

// Fetch User Friends
export async function fetchUserFriends(userId: string) {
  try {
    const result = await fetch(`${API}/user/friends?userId=${userId}`);
    const friends = await result.json();
    if (!friends?.length) {
      return [];
    }

    return friends;
  } catch (error) {
    console.error(error);
    return [];
  }
}

// Fetch all Users
export async function fetchUsers(userId: string) {
  try {
    const result = await fetch(`${API}/users/all/${userId}`);
    const users = await result.json();
    if (!users?.length) {
      return [];
    }

    return users;
  } catch (error) {
    console.error(error);
    return [];
  }
}

// Send a Friend Request to a User
export async function sendFriendRequest(userId: string, friendId: string) {
  try {
    const result = await fetch(`${API}/users/friends`, {
      method: "POST",
      body: JSON.stringify({ userId, friendId }),
    });
    if (!result.ok) {
      throw new Error("Failed to send friend request");
    }
  } catch (error) {
    throw new Error("Failed to send friend request: " + error);
  }
}

// Find a User by userId
export async function fetchUserById(userId: string) {
  try {
    const response = await fetch(`${API}/users?userId=${userId}`);
    const user = await response.json();
    const userData = {
      id: user.id,
      email: user.email,
      username: user.username,
      friendRequests: user.friendRequests,
    };

    return userData;
  } catch (error) {
    console.log(error);
    return {};
  }
}
