package gol

import (
	"fmt"
)

// type GolMessage struct {
// 	MessageType string `json:"msgType"`
// 	Payload string `json:"payload"`
// }

var ChatTestMessage = GolMessage{GolMsgType: "chat", Payload: "test chat payload"}
var GolGameTestMessage = GolMessage{GolMsgType: "GOLGAME", Payload: "test BOARD payload"}
var PColorTestMessage = GolMessage{GolMsgType: "playerColor", Payload: "test player color payload"}

//PrintBoard used for debugging
func PrintBoard(b *[][]Cell) {
	fmt.Printf("board is :\n %+v", *b)
}
