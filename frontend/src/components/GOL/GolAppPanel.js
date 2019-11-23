import React from "react";
import GolInput from "./GolInput";
import GolBoard from "./GolBoard";
import ChatHistory from "./GolChat/ChatHistory";
import ChatInput from "./GolChat/ChatInput";
import { Container, Row, Col } from "react-bootstrap";

export default function GolAppPanel() {
  return (
    <div>
      <Container>
        GOL MAIN Panel
        <hr />
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
