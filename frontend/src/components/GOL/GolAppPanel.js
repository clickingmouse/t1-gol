import React, { useEffect, useReducer } from "react";
import GolInput from "./GolInput";
import GolBoard from "./GolBoard";
import ChatHistory from "./GolChat/ChatHistory";
import ChatInput from "./GolChat/ChatInput";
import { Container, Row, Col } from "react-bootstrap";
import { connect, sendMsg } from "../../api";

function golMsgReducer(state, [type, payload]) {
  // console.log("REDUCER", type);
  // console.log("REDUCER-PAYLOAD", payload);
  let p = JSON.parse(payload.body);

  switch (p.msgType) {
    case "playerColor": {
      return Object.assign({}, state, {
        pending: state.pending,
        playerColor: p.colorHex,
        messageHistory: {
          chat: state.messageHistory.chat,
          game: state.messageHistory.game,
          allHistory: [...state.messageHistory.allHistory, payload]
        }
      });
    }
    case "chat":
      return Object.assign({}, state, {
        pending: state.pending,
        playerColor: state.playerColor,
        messageHistory: {
          chat: [...state.messageHistory.chat, p],
          game: [...state.messageHistory.game],
          allHistory: [...state.messageHistory.allHistory, payload]
        }
      });
    case "GOLGAME": {
      //console.log("game--------------->", p);
      return Object.assign({}, state, {
        pending: state.pending,
        playerColor: state.playerColor,
        messageHistory: {
          chat: [...state.messageHistory.chat],
          game: [...state.messageHistory.game, p],
          allHistory: [...state.messageHistory.allHistory, payload]
        }
      });
    }
    default:
      return state;
  }
}
const initialMsgState = {
  playerColor: "",
  pending: 0,
  messageHistory: { chat: [], game: [], allHistory: [] }
};
///////////////////////////////////////////////////////////////
//
//
//
///////////////////////////////////////////////////////////////
export default function GolAppPanel() {
  const [state, dispatch] = useReducer(golMsgReducer, initialMsgState);
  useEffect(() => {
    connect(msg => {
      console.log("New Message");
      console.log("received data:", typeof msg.data, msg.data);
      let packet = JSON.parse(msg.data);
      console.log("packet ->", packet);
      //dispatch([packet.type, packet]);
    });
  });

  //  connect();
  function send() {
    console.log("hello");
    sendMsg("hello");
  }
  return (
    <div>
      <Container>
        GOL MAIN Panel
        <hr />
        <button onClick={send}>Hit</button>
        <GolInput />
        <Row>
          <Col sm={7}>
            <GolBoard />
          </Col>
          <Col sm={5}>
            <ChatHistory />
            <ChatInput />
          </Col>
        </Row>
      </Container>
    </div>
  );
}
