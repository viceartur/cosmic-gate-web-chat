import { useState } from "react";

export default function Messsages(props: any) {
  const { socket } = props;
  const [messages, setMessages] = useState<String[]>([""]);
  const [message, setMessage] = useState<String>("");

  const handleSendMessage = () => {
    if (!props.socket) return;

    const obj = {
      type: "chat-message",
      receiverId: 2,
      message,
    };
    const jsonString = JSON.stringify(obj);

    socket.send(jsonString);
    setMessages([...messages, message]);
    setMessage("");
  };

  const onInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setMessage(e.target.value);
  };

  const sendMessage = () => {
    console.log("clicked");
    handleSendMessage();
  };

  return (
    <>
      <div className="messages-area">
        {messages.map((m: any, i: number) => (
          <p key={i}>{m}</p>
        ))}
      </div>
      <input
        type="text"
        onChange={onInputChange}
        placeholder="type any message"
      />
      <button onClick={sendMessage}>Send</button>
    </>
  );
}
