package pkg

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func RunClient(send chan string, ipAddr string) {
	conn, _, err := websocket.DefaultDialer.Dial(ipAddr, nil)

	if err != nil {
		log.Fatal("Unable to connect")
	}

	defer conn.Close()

	for {
		select {
		case msg := <-send:
			if conn == nil {
				fmt.Println("Not connected")
			}
			conn.WriteJSON(msg)
		}
	}
}
