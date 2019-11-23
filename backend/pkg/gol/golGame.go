package gol

type GameHandle struct {
	MessageType string `json:"msgType"`
	GameNum     int    `json:"gameNum"`
	//	Players     *[]Player `json:"players"`
	//	Board       *[][]Cell `json:"board"`
	GameMoveNum int    `json:"gameMoveNum"`
	GameID      string `json:"gameId"`
}
