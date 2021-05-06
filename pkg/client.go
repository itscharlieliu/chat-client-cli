package pkg

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func messageListener(conn *websocket.Conn) {
	for {
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			log.Panicln(err)
			return
		}
		// Return cursor to beginning of line
		fmt.Print("\033[2A\033[99C")
		fmt.Println("Recieve: " + string(bytes))
	}
}

func RunClient(send chan string, ipAddr string) {
	conn, _, err := websocket.DefaultDialer.Dial(ipAddr, nil)

	if err != nil {
		log.Fatal("Unable to connect")
	}

	defer conn.Close()

	go messageListener(conn)

	for {
		msg := <-send
		if conn == nil {
			fmt.Println("Not connected")
		}
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
	}

}
