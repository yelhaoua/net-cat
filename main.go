package main

import (
	"fmt"
	"os"

	"netcat/functions"
)

func main() {
	// defaul port declared
	port := ":8989"

	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if len(os.Args) == 1 {
		fmt.Println("Listening on the port :8989")
	} else {
		port = ":" + os.Args[1]
		fmt.Println("Listening on the port", port)
	}
	functions.Listiner(port)
}
