package main

import (
	"github.com/itscharlieliu/chat-client-cli/pkg"
)

func main() {
	pkg.ConnectWs("ws://127.0.0.1:8080")

}
