import React, { useEffect, useReducer } from "react";
import GolInput from "./GolInput";
import GolBoard from "./GolBoard";
import ChatHistory from "./GolChat/ChatHistory";
import ChatInput from "./GolChat/ChatInput";
import { Container, Row, Col } from "react-bootstrap";
import { connect, sendMsg } from "../../api";
// dispatch([body.golMsgType, body]);
function golMsgReducer(state, [type, body]) {
  // console.log("REDUCER", type);
  // console.log("REDUCER-PAYLOAD", payload);
  //let p = JSON.parse(payload.body);
  let p = { msgType: type };
  //let p = payload;
  switch (p.msgType) {
    case "playerColor": {
      return Object.assign({}, state, {
        pending: state.pending,
        playerColor: body.payload,
        messageHistory: {
          chat: state.messageHistory.chat,
          game: state.messageHistory.game,
          allHistory: [...state.messageHistory.allHistory, body.payload]
        }
      });
    }
    case "chat":
      return Object.assign({}, state, {
        pending: state.pending,
        playerColor: state.playerColor,
        messageHistory: {
          chat: [...state.messageHistory.chat, body.payload],
          game: [...state.messageHistory.game],
          allHistory: [...state.messageHistory.allHistory, body.payload]
        }
      });
    case "GOLGAME": {
      //console.log("game--------------->", p);
      return Object.assign({}, state, {
        pending: state.pending,
        playerColor: state.playerColor,
        messageHistory: {
          chat: [...state.messageHistory.chat],
          game: [...state.messageHistory.game, body.payload],
          allHistory: [...state.messageHistory.allHistory, body.payload]
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
      //console.log("New Message");
      //console.log("received data:", typeof msg.data, msg.data);
      let packet = JSON.parse(msg.data);
      //console.log("packet ->", packet.golMsgType, packet);
      //console.log("body ->", typeof packet.body, packet.body);
      let body = JSON.parse(packet.body);
      //console.log(typeof payload, payload.golMsgType);

      //      dispatch([packet.golMsgType, packet]);
      dispatch([body.golMsgType, body]);
    });
  }, []);

  //  connect();
  function send(data) {
    console.log("sending...:", data);
    sendMsg(data);
  }

  //   <GolBoard
  //   gameHistory={state.messageHistory.game}
  //   myColor={state.playerColor}
  // />
  console.log("STATE", state);
  // {state.messageHistory.game.length >= 1 ? (
  //   <GolBoard
  //     gameHistory={state.messageHistory.game}
  //     myColor={state.playerColor}
  //   />
  // ) : null}

  return (
    <div>
      <Container>
        GOL MAIN Panel log::{state.messageHistory.allHistory.length}
        <hr />
        <button onClick={send}>Hit</button>
        <GolInput send={send} />
        <Row>
          <Col sm={7}>
            {state.messageHistory.game.length >= 1 ? (
              <GolBoard
                boardData={
                  JSON.parse(
                    state.messageHistory.game[
                      state.messageHistory.game.length - 1
                    ]
                  ).board
                }
                myColor={state.playerColor}
                send={send}
              />
            ) : null}
          </Col>
          <Col sm={5}>
            <ChatHistory
              chatHistory={state.messageHistory.chat}
              header={"Chat Only"}
            />
            <ChatInput send={send} />
          </Col>
        </Row>
      </Container>
    </div>
  );
}
