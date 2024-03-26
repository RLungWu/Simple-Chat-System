package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//Define upgrader
//We need and Read and Write Buffer
var upgrader = websocket.Upgrader{
	ReadBufferSize : 1024,
	WriteBufferSize: 1024,

	//Check the origin of connection
	//This will allow our request from our frontend
	//But now, we just allow the connection
	CheckOrigin: func(r *http.Request) bool {return true},
}

//Define reader to listen
//New message will be sent to our WebSocket point
func reader(conn *websocket.Conn){
	for{
		messageType, p, err := conn.ReadMessage()
		
	}
}
