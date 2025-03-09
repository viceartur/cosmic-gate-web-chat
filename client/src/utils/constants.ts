export const API: string = "http://localhost:8080";

export interface Message {
  id: string;
  senderId: number;
  recipientId: number;
  text: string;
  sentAt: Date;
}
