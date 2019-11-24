import React, { useState } from "react";

export default function ChatInput(props) {
  const [input, setInput] = useState("");

  const handleChange = e => {
    setInput(e.target.value);
  };
  const send = e => {
    //    console.log("hello");
    //   sendMsg("hello");
    if (e.keyCode === 13) {
      const pMsg = {
        msgType: "GOLCHAT",
        X: "",
        Y: "",
        playerColor: props.myColor,
        generation: 99,
        playerID: "007",
        payload: input
      };
      const message = {
        msgType: "GOLCHAT",
        //msgData: e.target.value
        payload: input
      };
      //++sendMsg(JSON.stringify(message));
      //sendMsg(e.target.value);
      console.log("sending", pMsg);
      props.send(JSON.stringify(pMsg));
      e.target.value = "";
      console.log("clearing");
      setInput("");
    }
  };

  return (
    <div>
      <div className="ChatInput">
        <input onKeyDown={send} onChange={handleChange} value={input} />
      </div>
    </div>
  );
}
