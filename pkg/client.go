package pkg

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Send chan string
	Exit chan bool
}

func messageListener(conn *websocket.Conn) {
	for {
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			log.Panicln(err)
			return
		}
		fmt.Println("Recieve: " + string(bytes))
	}
}

func RunClient(client Client, ipAddr string, listen bool, wg *sync.WaitGroup) {

	conn, _, err := websocket.DefaultDialer.Dial(ipAddr, nil)

	if err != nil {
		log.Fatal("Unable to connect")
	}

	defer conn.Close()
	defer wg.Done()

	if listen {
		go messageListener(conn)
	}

	for {
		select {
		case msg := <-client.Send:
			if conn == nil {
				fmt.Println("Not connected")
			}
			conn.WriteMessage(websocket.TextMessage, []byte(msg))
		case <-client.Exit:
			return
		}

	}
}
