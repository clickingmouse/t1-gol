package gol

import (
	"fmt"
	"time"
	//"github.com/clickingmouse/t1/gol/pkg/websocket"
)

type GolMessage struct {
	GolMsgType string `json:"golMsgType"`
	Payload    string `json:"payload"`
}

type GolTextWrapper struct {
	Type int        `json:"type"`
	Body GolMessage `json:"body"`
}

type GolGameMessage struct {
	GolMsgType string     `json:"golMsgType"`
	Payload    GameHandle `json:"payload"`
}

type GolGameWrapper struct {
	Type int            `json:"type"`
	Body GolGameMessage `json:"body"`
}

func NewMsg(msgType, payload string) *GolMessage {
	return &GolMessage{GolMsgType: msgType, Payload: payload}
}

//
var NewUsrMsg = GolTextWrapper{Type: 1, Body: GolMessage{GolMsgType: "chat", Payload: "A New Gamer Joined!"}}

//var NewUsrMsg = GolTextWrapper{GolMsgType: "chat", Payload: "New User Joined"}

type GolPlayerMsg struct {
	MsgType    string `json:"msgType"`     //: "GOLGAME",
	X          int    `json:"X"`           //,
	Y          int    `json:"Y"`           //,
	Color      string `json:"playerColor"` //: props.myColor,
	Generation int    `json:"generation"`  // 99,
	PlayerID   string `json:"playerID"`    //: "007",
	//moveType: "instill",
	//text: "",
	Payload string `json:"payload"` //: "instill"

}

////////////////////////////////////////////////////////////////////////
//
//
//
func PlayerAction(pMsg *GolPlayerMsg, gH GameHandle, t *time.Ticker) string {
	fmt.Printf("PlayerAction:: rec'd pMsg: %+v\n", pMsg)
	switch pMsg.MsgType {
	case "GOLMOVE":
		fmt.Printf("MAKE [%d][%d]CELL ALIVE!!\n", pMsg.X, pMsg.Y)
		//func (c *Cell) Breed(color string)
		(*gH.Board)[pMsg.Y][pMsg.X].Breed(pMsg.Color)
		return "moved"
	case "GOLCHAT":
		fmt.Printf("CHAT MESSAGE :%s\n", pMsg.Payload)
		return pMsg.Payload
	case "ANNIHILATE":
		fmt.Printf("RESETTING ... ... .. .\n")
		BoardClearAll(gH.Board)
		// reset timer too
		// seems there are channels issue, will have to look into this in the future
		//t.Stop()
		//ticker := time.NewTicker(5000 * time.Millisecond)
		//t = time.NewTicker(1000 * time.Millisecond)
		return "RESET"

	case "PROPOGATE":
		fmt.Printf("PROPAGATING ... ... .. .\n")
		Propagate(gH.Board)
		return "NEW GENERATION"

	case "BLINKER":
		fmt.Printf("LOAD BLINKER ... ... .. .\n")
		BoardClearAll(gH.Board)
		PreloadBlinker(gH.Board, pMsg.Color)
		return "BLINKER"
	case "TOAD":
		fmt.Printf("LOAD TOAD ... ... .. .\n")
		BoardClearAll(gH.Board)
		PreloadToad(gH.Board, pMsg.Color)
		return "TOAD"
	case "BEACON":
		fmt.Printf("LOAD BEACON ... ... .. .\n")
		BoardClearAll(gH.Board)
		PreloadBeacon(gH.Board, pMsg.Color)
		return "BEACON"
	default:
		return "nothing"
	}

	return "something"
}
