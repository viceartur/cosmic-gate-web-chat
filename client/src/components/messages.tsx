import { useEffect, useState } from "react";

import { fetchMessages } from "@/actions/messages";

export default function Messages(props: any) {
  const { socket, userId, friendId } = props;
  const [messages, setMessages] = useState<
    { senderId: string; data: string }[]
  >([]);
  const [message, setMessage] = useState<object | any>({
    senderId: "",
    data: "",
  });

  // Get Message History for Users
  useEffect(() => {
    const getMessages = async () => {
      const messages = await fetchMessages(userId, friendId);
      setMessages(messages);
    };
    getMessages();
  }, [userId]);

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
      senderId: userId, // auth ID
      recipientId: friendId, // friend ID
      data: message.data,
    };
    const jsonString = JSON.stringify(obj);

    socket.send(jsonString);

    setMessages((prevMessages) => [...prevMessages, message]);
    setMessage({ senderId: "", data: "" });
  };

  const onInputChange = (e: any) => {
    setMessage({ senderId: userId, data: e.target.value });
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
              m.senderId === userId
                ? "messages__message--sent"
                : "messages__message--received"
            }`}
          >
            {m.data}
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
