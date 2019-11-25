import React from "react";

export default function ChatHistory(props) {
  const messages = props.chatHistory.map((msg, index) => {
    //<Message message={msg.body} />
    //console.log(msg);
    return <li key={index}>{JSON.stringify(msg)}</li>;
  });
  return (
    <div
      className="ChatHistory"
      style={{ height: "70vh", overflowY: "scroll" }}
    >
      <h2>Chat History</h2>
      <h3>{props.header}</h3>
      <ul>{messages}</ul>
    </div>
  );
}
