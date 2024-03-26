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
		if err != nil{
			log.Println(err)
			return
		}

		//Log out the message
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil{
			log.Println(err)
			return
		}
	}

}


//Define Websocket Points
func serveWs(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Host)

	//Upgrade to Websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(err)
	}

	//Liten to new message and connect with our websocket
	reader(ws)
}

func setupRoutes(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Simple Server")
	})

	//Connect '/ws' to serveWs
	http.HandleFunc("/ws", serveWs)
}

func main(){
	fmt.Println("Chat App")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
