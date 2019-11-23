package gol

type GolMessage struct {
	GolMsgType string `json:"golMsgType"`
	Payload    string `json:"payload"`
}

func NewMsg(msgType, payload string) *GolMessage {
	return &GolMessage{GolMsgType: msgType, Payload: payload}

}
