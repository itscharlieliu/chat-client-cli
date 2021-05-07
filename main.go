package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/itscharlieliu/chat-client-cli/pkg"
)

func main() {
	listen := flag.Bool("l", false, "listen")

	flag.Parse()

	client := pkg.Client{
		Send: make(chan string),
		Exit: make(chan bool),
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go pkg.RunClient(client, "ws://127.0.0.1:8080", *listen, &wg)

	// We don't want to send messages if we are listening
	if !*listen {
		for {
			fmt.Print("Send: ")
			var msg string
			fmt.Scanln(&msg)

			client.Send <- msg
		}
	} else {
		fmt.Println("Listening...")
	}

	// Handle sigtern interrupts
	end := make(chan os.Signal)
	signal.Notify(end, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-end
		client.Exit <- true
	}()

	wg.Wait()
}
