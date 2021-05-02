package pkg

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func ConnectWs(addr string) {
	_, resp, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Body)
}
