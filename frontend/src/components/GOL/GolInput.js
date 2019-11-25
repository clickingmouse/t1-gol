import React from "react";
import { Button, Row } from "react-bootstrap";
export default function GolInput(props) {
  console.log(props);
  const handleClick = e => {
    //dosomething
    const pClick = {
      msgType: e.target.id,
      X: null,
      Y: null,
      playerColor: props.myColor,
      generation: 99,
      playerID: "007",
      payload: e.target.id
    };
    console.log("sending", pClick);
    props.send(JSON.stringify(pClick));
  };
  return (
    <div>
      GOL OPTIONS
      <br />
      <Row
        style={{
          display: "flex",
          flex: 1,
          flexDirection: "row",
          justifyContent: "space-between",
          flexWrap: "no-wrap",
          //margin: "auto",
          padding: "1rem"
        }}
      >
        <Button variant="primary" id="PROPOGATE" onClick={handleClick}>
          PROPOGATE
        </Button>
        <Button variant="primary" id="BLINKER" onClick={handleClick}>
          BLINKER
        </Button>
        <Button variant="primary" id="TOAD" onClick={handleClick}>
          TOAD
        </Button>
        <Button variant="primary" id="BEACON" onClick={handleClick}>
          BEACON
        </Button>
        <Button variant="primary" id="ANNIHILATE" onClick={handleClick}>
          ANNIHILATE
        </Button>
      </Row>
      <hr />
    </div>
  );
}
