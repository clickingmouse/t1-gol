package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/clickingmouse/t1/gol/pkg/gol"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(reflect.TypeOf(p))
		message := Message{Type: messageType, Body: string(p)}
		_ = message
		// do somethine with message here/////////////////////
		var pMsg gol.GolPlayerMsg
		err = json.Unmarshal(p, &pMsg)

		if err != nil {
			log.Println("client.go E", err)
		}
		///////////////////////////////////////////////
		fmt.Printf("GOT PLAYER MSG:%+v\n", pMsg)
		// pass timer too
		playerMessage := gol.PlayerAction(&pMsg, *c.Pool.GameHandle, c.Pool.Ticker)
		//		c.Pool.Broadcast <- message
		c.Pool.Broadcast <- gol.GolTextWrapper{Type: 1, Body: gol.GolMessage{GolMsgType: "chat", Payload: playerMessage}}

		c.Pool.UpdateBoard <- (*c.Pool).GameHandle

		//fmt.Printf("Message Received: %+v\n", message)
	}
}
