package websocket

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/clickingmouse/t1/gol/pkg/gol"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	//Broadcast   chan Message
	Broadcast   chan gol.GolTextWrapper
	UpdateBoard chan *gol.GameHandle
	Timer       chan *time.Ticker
	GameHandle  *gol.GameHandle
	Ticker      *time.Ticker
	//	Tm          *time
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		//Broadcast:   make(chan Message),
		Broadcast: make(chan gol.GolTextWrapper),

		UpdateBoard: make(chan *gol.GameHandle),
		GameHandle:  gol.InitNewGame(6, 6),
	}
}

//4ref; func InitNewGame(r, c int) so  7 rows, 9 cols

func (pool *Pool) Start() {
	// non empty board for frontend testing
	//gol.InsertDummyData(pool.GameHandle.Board)
	//gol.PrintBoard(pool.GameHandle.Board)

	pool.Ticker = time.NewTicker(5000 * time.Millisecond)
	//pool.Tm.Ticker = time.NewTicker(5000 * time.Millisecond)
	for {
		select {
		case Timer := <-pool.Ticker.C:
			fmt.Println("Tick at", Timer)
			gol.Propagate(pool.GameHandle.Board)
			for client, _ := range pool.Clients {

				_, err := json.Marshal(*pool.GameHandle)
				if err != nil {
					panic(err)
				}
				(*client).Conn.WriteJSON(gol.GolGameWrapper{Type: 1, Body: gol.GolGameMessage{GolMsgType: "GOLGAME", Payload: *pool.GameHandle}})

			}
			// propagate here & broadcast
			//pool.UpdateBoard <- pool.GameHandle
			break
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			// push random color to client
			//create new message obj
			var RandomColorMessage = gol.GolMessage{GolMsgType: "playerColor", Payload: gol.GetRandomColor()}
			_, err := json.Marshal(&RandomColorMessage)
			if err != nil {
				panic(err)
			}
			//OLD			(*client).Conn.WriteJSON(Message{Type: 1, Body: string(nPC)})
			(*client).Conn.WriteJSON(gol.GolTextWrapper{Type: 1, Body: RandomColorMessage})

			for client, _ := range pool.Clients {
				fmt.Println(client)

				// send new user joined message
				_, err := json.Marshal(&gol.NewUsrMsg)
				if err != nil {
					panic(err)
				}

				(*client).Conn.WriteJSON(gol.NewUsrMsg)

				// send Game & more importantly, game board
				_, err = json.Marshal(*pool.GameHandle)
				if err != nil {
					panic(err)
				}
				(*client).Conn.WriteJSON(gol.GolGameWrapper{Type: 11, Body: gol.GolGameMessage{GolMsgType: "GOLGAME", Payload: *pool.GameHandle}})

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

				_, err := json.Marshal(*pool.GameHandle)
				if err != nil {
					panic(err)
				}
				(*client).Conn.WriteJSON(gol.GolGameWrapper{Type: 1, Body: gol.GolGameMessage{GolMsgType: "GOLGAME", Payload: *pool.GameHandle}})

			}
			break
		}
	}
}
