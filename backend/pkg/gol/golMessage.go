package gol

import "fmt"

type GolMessage struct {
	GolMsgType string `json:"golMsgType"`
	Payload    string `json:"payload"`
}

func NewMsg(msgType, payload string) *GolMessage {
	return &GolMessage{GolMsgType: msgType, Payload: payload}
}

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

func PlayerAction(pMsg *GolPlayerMsg, gH GameHandle) string {
	fmt.Printf("PlayerAction:: rec'd pMsg: %+v", pMsg)
	switch pMsg.MsgType {
	case "GOLMOVE":
		fmt.Printf("MAKE [%d][%d]CELL ALIVE!!\n", pMsg.X, pMsg.Y)
		//func (c *Cell) Breed(color string)
		(*gH.Board)[pMsg.Y][pMsg.X].Breed(pMsg.Color)
		return "moved"
	case "GOLCHAT":
		fmt.Printf("CHAT MESSAGE :%s\n", pMsg.Payload)
		return pMsg.Payload

	case "PROPOGATE":
		fmt.Printf("PROPAGATING ... ... .. .\n")
		Propagate(gH.Board)
	default:
		return "nothing"
	}

	return "something"
}
