package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	log.Println("Websocket Server Started!")
	wsserver := socketio.NewServer(nil)

	wsserver.OnConnect("/", func(c socketio.Conn) error {
		c.SetContext("")
		fmt.Println("connected : ", c.ID())
		return nil
	})

	http.ListenAndServe(":5000", nil)
}
