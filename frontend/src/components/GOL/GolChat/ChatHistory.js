import React from "react";
import { ListGroup } from "react-bootstrap";
export default function ChatHistory(props) {
  const messages = props.chatHistory.map((msg, index) => {
    //<Message message={msg.body} />
    //console.log(msg);
    return <li key={index}>{"@:" + JSON.stringify(msg)}</li>;
  });
  //  <h2>Chat History</h2>
  return (
    <div
      className="ChatHistory border"
      style={{ height: "70vh", overflowY: "scroll" }}
    >
      <h3>{props.header}</h3>
      <ul style={{ listStyleType: "none" }}>{messages}</ul>
    </div>
  );
}
