package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/clickingmouse/t1/gol/pkg/websocket"
)

//func serveWs(w http.ResponseWriter, r *http.Request) {
// ws, err := websocket.Upgrade(w, r)
// if err != nil {
// 	fmt.Fprintf(w, "%+V\n", err)
// }
// go websocket.Writer(ws)
// websocket.Reader(ws)
//}
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	//http.HandleFunc("/ws", serveWs)
	// TODO: add gol game initialization inside newPool
	pool := websocket.NewPool()
	fmt.Printf("%+v", pool.GameHandle)
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

	http.Handle("/", http.FileServer(http.Dir("./web")))

	//dir, _ := os.Getwd()
	// fs := http.FileServer(http.Dir(dir))
	// http.Handle("/", fs)
	// http.HandleFunc("/gol/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./build/index.html")
	// })
	//log.Println("Serving " + dir)

}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080", nil

		//return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	fmt.Println("Distributed T1 - GOL v0.01")
	setupRoutes()
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
