"use server";

import { API, Message } from "@/utils/constants";

// Fetch messages between two users
export async function fetchMessages(senderId: string, recipientId: string) {
  try {
    const queryParams = new URLSearchParams({
      senderId,
      recipientId,
    });
    const result = await fetch(`${API}/messages?${queryParams.toString()}`);
    const data = await result.json();
    if (!data?.length) {
      return [];
    }
    const messages = data.map((message: Message) => ({
      senderId: message.senderId,
      data: message.text,
    }));
    return messages;
  } catch (error) {
    console.error(error);
    return [];
  }
}
