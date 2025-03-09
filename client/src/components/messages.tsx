import { useEffect, useState } from "react";

import { fetchMessages } from "@/actions/messages";

export default function Messsages(props: any) {
  const { socket, userId, friendId } = props;
  const [messages, setMessages] = useState<
    { senderId: number; data: string }[]
  >([]);
  const [message, setMessage] = useState<object | any>({
    senderId: 0,
    data: "",
  });

  // Get Message History for Users
  useEffect(() => {
    const getMessages = async () => {
      const messages = await fetchMessages(userId, friendId);
      setMessages(messages);
    };
    getMessages();
  }, []);

  // WebSocket Connection and Handle Messages
  useEffect(() => {
    if (socket) {
      socket.onmessage = (event: MessageEvent) => {
        const socketMsg = JSON.parse(event.data);
        const msg = {
          senderId: socketMsg.senderId,
          data: socketMsg.data,
        };

        if (socketMsg.type === "chat-connection") {
          setMessages((prevMessages) => [...prevMessages, msg]);
        } else if (socketMsg.type === "chat-message") {
          setMessages((prevMessages) => [...prevMessages, msg]);
        }
      };
    }
  }, [socket]);

  // Send Message to the Server via WebSocket
  const handleSendMessage = () => {
    if (!socket) return;

    const obj = {
      type: "chat-message",
      senderId: Number(userId), // auth ID
      recipientId: Number(friendId), // friend ID
      data: message.data,
    };
    const jsonString = JSON.stringify(obj);

    socket.send(jsonString);

    setMessages((prevMessages) => [...prevMessages, message]);
    setMessage({ senderId: 0, data: "" });
  };

  const onInputChange = (e: any) => {
    setMessage({ senderId: Number(userId), data: e.target.value });
  };

  const sendMessage = () => {
    handleSendMessage();
  };

  return (
    <div className="messages">
      <div className="messages__area">
        {messages.map((m: any, i: number) => (
          <div
            key={i}
            className={`messages__message ${
              m.senderId === Number(userId)
                ? "messages__message--sent"
                : "messages__message--received"
            }`}
          >
            {m.senderId}: {m.data}
          </div>
        ))}
      </div>
      <div className="messages__actions">
        <input
          className="messages__input"
          type="text"
          onChange={onInputChange}
          placeholder="type any message"
          value={message.data}
        />
        <button
          className="messages__button"
          type="button"
          onClick={sendMessage}
        >
          Send
        </button>
      </div>
    </div>
  );
}
