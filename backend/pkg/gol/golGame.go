package gol

import "github.com/google/uuid"

type GameHandle struct {
	//MessageType string `json:"msgType"`
	GameCount   int       `json:"gameNum"`
	Players     *[]Player `json:"players"`
	Board       *[][]Cell `json:"board"`
	GameMoveNum int       `json:"gameMoveNum"`
	GameID      string    `json:"gameId"`
}

func InitNewGame(r, c int) *GameHandle {
	gHandle := &GameHandle{}
	//gHandle.MessageType = "GOLGAME"
	gHandle.Board = CreateNewBoard(r, c)
	gHandle.Board = InitBoard(gHandle.Board)

	//should add err checking too
	id, err := uuid.NewUUID()
	if err != nil {
		// handle error
	}
	gHandle.GameID = id.String()

	// set player to 1 for now
	gPlayer := Player{"007", "#ff0000"}
	gHandle.Players = &[]Player{gPlayer}

	return gHandle
}
