package websocket

import (
    "fmt"
    "io"
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
/*
Why we have to upgrade to websocket?
Real Time Interaction
Efficiency and Lower Costs
Better Use of Server Resources
*/

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(err)
		return conn, err
	}

	return conn, nil
}


//Define reader to listen
//New message will be sent to our WebSocket point
func Reader(conn *websocket.Conn) {
    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }

        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }
    }
}



func Writer(conn *websocket.Conn) {
    for {
        fmt.Println("Sending")

        messageType, r, err := conn.NextReader()
        if err != nil {
            fmt.Println(err)
            return
        }

        w, err := conn.NextWriter(messageType)
        if err != nil {
            fmt.Println(err)
            return
        }

        if _, err := io.Copy(w, r); err != nil {
            fmt.Println(err)
            return
        }

        if err := w.Close(); err != nil {
            fmt.Println(err)
            return
        }
    }
}