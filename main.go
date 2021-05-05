package main

import (
	"fmt"

	"github.com/itscharlieliu/chat-client-cli/pkg"
)

func main() {
	send := make(chan string)

	go pkg.RunClient(send, "ws://127.0.0.1:8080")

	for {
		fmt.Print("Send: ")
		var msg string
		fmt.Scanln(&msg)
		fmt.Println(msg)
		send <- msg
	}

}
