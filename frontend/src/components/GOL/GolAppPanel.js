import React, { useEffect, useReducer } from "react";
import GolInput from "./GolInput";
import GolBoard from "./GolBoard";
import ChatHistory from "./GolChat/ChatHistory";
import ChatInput from "./GolChat/ChatInput";
import { Container, Row, Col } from "react-bootstrap";
import { connect, sendMsg } from "../../api";
// dispatch([body.golMsgType, body]);
//packet.body.payload.golMsgType
function golMsgReducer(state, [type, body]) {
  console.log("REDUCER", type);
  console.log("REDUCER-PAYLOAD", body);
  //let p = JSON.parse(payload.body);
  //let p = { msgType: type };
  //let p = payload;
  switch (type) {
    case "playerColor": {
      return Object.assign({}, state, {
        pending: state.pending,
        playerColor: body,
        messageHistory: {
          chat: state.messageHistory.chat,
          game: state.messageHistory.game,
          allHistory: [...state.messageHistory.allHistory, body]
        }
      });
    }
    case "chat":
      return Object.assign({}, state, {
        pending: state.pending,
        playerColor: state.playerColor,
        messageHistory: {
          chat: [...state.messageHistory.chat, body],
          game: [...state.messageHistory.game],
          allHistory: [...state.messageHistory.allHistory, body]
        }
      });
    case "GOLGAME": {
      //console.log("game--------------->", p);
      return Object.assign({}, state, {
        pending: state.pending,
        playerColor: state.playerColor,
        messageHistory: {
          chat: [...state.messageHistory.chat],
          game: [...state.messageHistory.game, body],
          allHistory: [...state.messageHistory.allHistory, body]
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
      console.log("packet ->", packet.type, packet);

      //body is an object
      console.log("body ->", typeof packet.body, packet.body);
      console.log("body ->", packet.body.golMsgType, packet.body.payload);

      //
      //
      //let body = JSON.parse(packet.body);
      //console.log("++++++++", packet.body.golMsgType);

      //      dispatch([packet.golMsgType, packet]);
      dispatch([packet.body.golMsgType, packet.body.payload]);
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
        <GolInput send={send} myColor={state.playerColor} />
        <Row>
          <Col sm={7}>
            {state.messageHistory.game.length >= 1 ? (
              <GolBoard
                boardData={
                  state.messageHistory.game[
                    state.messageHistory.game.length - 1
                  ].board
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
