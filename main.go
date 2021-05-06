package main

import (
	"fmt"

	"github.com/itscharlieliu/chat-client-cli/pkg"
)

func main() {
	send := make(chan string)

	go pkg.RunClient(send, "ws://127.0.0.1:8080")

	/*
		We are using escape codes ("\033[") to modify the terminal
		2J - Clear the screen
		0;0H - Move cursor to line 0, column 0
	*/
	fmt.Print("\033[2J\033[3;0H")

	for {
		fmt.Print("Send: ")
		var msg string
		fmt.Scanln(&msg)

		send <- msg
	}
}
