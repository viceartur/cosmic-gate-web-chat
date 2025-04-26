import { API } from "@/utils/constants";

// Fetch User Friends
export async function fetchUserFriends(userId: string) {
  try {
    const result = await fetch(`${API}/users/friends?userId=${userId}`);
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
    const result = await fetch(`${API}/users/friend-requests`, {
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
      bio: user.bio,
    };

    return userData;
  } catch (error) {
    console.log(error);
    return {};
  }
}

// Fetch User Friend Requests
export async function fetchUserFriendRequests(userId: string) {
  try {
    const result = await fetch(`${API}/users/friend-requests?userId=${userId}`);
    const friendRequestSenders = await result.json();
    if (!friendRequestSenders?.length) {
      return [];
    }

    return friendRequestSenders;
  } catch (error) {
    console.error(error);
    return [];
  }
}

// Accept a Friend Request
export async function acceptFriendRequest(userId: string, friendId: string) {
  try {
    const result = await fetch(`${API}/friend-request/accept`, {
      method: "POST",
      body: JSON.stringify({ userId, friendId }),
    });
    if (!result.ok) {
      throw new Error("Failed to accept a friend request");
    }
  } catch (error) {
    throw new Error("Failed to accept a friend request: " + error);
  }
}

// Decline a Friend Request
export async function declineFriendRequest(userId: string, friendId: string) {
  try {
    const result = await fetch(`${API}/friend-request/decline`, {
      method: "POST",
      body: JSON.stringify({ userId, friendId }),
    });
    if (!result.ok) {
      throw new Error("Failed to decline a friend request");
    }
  } catch (error) {
    throw new Error("Failed to decline a friend request: " + error);
  }
}

// Update User Data
export async function updateUserData(userId: string, formData: FormData) {
  try {
    const result = await fetch(`${API}/users`, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id: userId,
        email: formData.get("email"),
        username: formData.get("username"),
        bio: formData.get("bio"),
        password: formData.get("password"),
      }),
    });

    const data = await result.json();

    if (!result.ok) {
      throw new Error(data.error || `Server Error: ${result.status}`);
    }

    return data;
  } catch (error: any) {
    throw new Error(error.message);
  }
}
