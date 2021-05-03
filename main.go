package main

import (
	"strconv"
	"time"

	"github.com/itscharlieliu/chat-client-cli/pkg"
)

func main() {
	send := make(chan string)

	go pkg.RunClient(send, "ws://127.0.0.1:8080")

	for i := 0; true; i++ {
		send <- "test" + strconv.Itoa(i)
		time.Sleep(1 * time.Second)
	}

}
