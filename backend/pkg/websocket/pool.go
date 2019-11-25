package websocket

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/clickingmouse/t1/gol/pkg/gol"
)

type Pool struct {
	Register    chan *Client
	Unregister  chan *Client
	Clients     map[*Client]bool
	Broadcast   chan Message
	UpdateBoard chan *gol.GameHandle
	Timer       chan *time.Ticker
	GameHandle  *gol.GameHandle
	Ticker      *time.Ticker
	//	Tm          *time
}

func NewPool() *Pool {
	return &Pool{
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		Clients:     make(map[*Client]bool),
		Broadcast:   make(chan Message),
		UpdateBoard: make(chan *gol.GameHandle),
		GameHandle:  gol.InitNewGame(6, 6),
	}
}

//4ref; func InitNewGame(r, c int) so  7 rows, 9 cols

func (pool *Pool) Start() {
	// non empty board for frontend testing
	gol.InsertDummyData(pool.GameHandle.Board)
	//gol.PrintBoard(pool.GameHandle.Board)

	pool.Ticker = time.NewTicker(5000 * time.Millisecond)
	//pool.Tm.Ticker = time.NewTicker(5000 * time.Millisecond)
	for {
		select {
		case Timer := <-pool.Ticker.C:
			fmt.Println("Tick at", Timer)
		// propagate here
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			// push random color to client
			//create new message obj
			var RandomColorMessage = gol.GolMessage{GolMsgType: "playerColor", Payload: gol.GetRandomColor()}
			nPC, err := json.Marshal(&RandomColorMessage)
			if err != nil {
				panic(err)
			}
			(*client).Conn.WriteJSON(Message{Type: 1, Body: string(nPC)})

			for client, _ := range pool.Clients {
				fmt.Println(client)
				//client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
				// test messages

				//var NewUserMessage = GolMessage{GolMsgType: "chat", Payload: "New User Joined"}
				nUsrMsg, err := json.Marshal(&gol.NewUserMessage)
				if err != nil {
					panic(err)
				}
				(*client).Conn.WriteJSON(Message{Type: 1, Body: string(nUsrMsg)})

				//client.Conn.WriteJSON(gol.NewUserMessage)
				//client.Conn.WriteJSON(gol.ChatTestMessage)
				//client.Conn.WriteJSON(gol.PColorTestMessage)
				//client.Conn.WriteJSON(gol.GolGameTestMessage)

				// send Game & more importantly, game board
				pGH, err := json.Marshal(*pool.GameHandle)
				if err != nil {
					panic(err)
				}
				//gBoardMessage := gol.NewMsg("GOLGAME", string(pGH))
				//client.Conn.WriteJSON(gBoardMessage)
				//-(*client).Conn.WriteJSON(Message{Type: 1, Body: string(pGH)})

				//////////////////////////////////////////////////////////////////
				var boardStatusMsg = gol.GolMessage{GolMsgType: "GOLGAME", Payload: string(pGH)}
				bStatus, err := json.Marshal(&boardStatusMsg)
				if err != nil {
					panic(err)
				}
				(*client).Conn.WriteJSON(Message{Type: 1, Body: string(bStatus)})

			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {

				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		case updateBoard := <-pool.UpdateBoard:
			fmt.Println("Updating board to all clients in Pool %+v\n", updateBoard)
			for client, _ := range pool.Clients {
				pGH, err := json.Marshal(*pool.GameHandle)
				if err != nil {
					panic(err)
				}
				var boardStatusMsg = gol.GolMessage{GolMsgType: "GOLGAME", Payload: string(pGH)}
				bStatus, err := json.Marshal(&boardStatusMsg)
				if err != nil {
					panic(err)
				}
				(*client).Conn.WriteJSON(Message{Type: 1, Body: string(bStatus)})
			}
			break
		}
	}
}
