package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RLungWu/Simple-Chat-System/pkg/websocket"
)

// Define Websocket Points
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")

	//Upgrade to Websocket connection
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprint(w, "%+V\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	/* Test Over
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Simple Server")
	})
	*/

	pool := websocket.NewPool()

	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
