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
