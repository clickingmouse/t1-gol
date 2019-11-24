package gol

import "fmt"

type GolMessage struct {
	GolMsgType string `json:"golMsgType"`
	Payload    string `json:"payload"`
}

type GolTextWrapper struct {
	Type string     `json:"golMsgType"`
	Body GolMessage `json:"body"`
}

type GolGameMessage struct {
	GolMsgType string     `json:"goMsgType`
	Payload    GameHandle `json:"payload"`
}

type GolGameWrapper struct {
	Type string `json:"golMsgType"`
	Body string `json:"body"`
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
		break
	case "RESET":
		fmt.Printf("RESETING ... ... .. .\n")
		return "RESET"
	case "BLINKER":
		fmt.Printf("LOAD BLINKER ... ... .. .\n")
		return "BLINKER"
	case "TOAD":
		fmt.Printf("LOAD TOAD ... ... .. .\n")
		return "TOAD"
	case "BEACON":
		fmt.Printf("LOAD BEACON ... ... .. .\n")
		return "BEACON"
	default:
		return "nothing"
	}

	return "something"
}
