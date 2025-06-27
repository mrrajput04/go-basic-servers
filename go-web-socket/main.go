package main

import (
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/net/websocket"
)

// the server will listen for incoming connections on this port
var address = ":4321"

// takes a pointer to a `websocket.Conn` object as input.
// This object represents the established connection with the client
func LengthServerResponse(ws *websocket.Conn) {

	//this variable stores messages received from the client
	var msg string

	// this keeps running as long as the connection is open.
	// what determines an open connection ?

	for {
		//we receive the message from the client and store in `msg`
		websocket.Message.Receive(ws, &msg)
		fmt.Println("msg got", msg)

		//calculates the length of the received message
		msgLen := len(msg)

		//converts int length to string, sens the string back through established websocket connection
		//checks for errors while processing

		if err := websocket.Message.Send(ws, strconv.FormatInt(int64(msgLen), 10)); err != nil {
			fmt.Println("cannot send message length")
			break
		}
	}
}

func websocketListen() {
	//when a client connects to the `/length` path, the `websocket.Handler` uses
	//the logic in function created earlier to handle WebSocket communication.
	http.Handle("/length", websocket.Handler(LengthServerResponse))
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		//this serves the html file to the client as a response to `/websocket` request
		http.ServeFile(w, r, "index.html")
	})
	websocketListen()
}
