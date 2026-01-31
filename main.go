package main

import (
	"fmt"
	"net-cat/handlers"
	"os"
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
	}
	if len(os.Args) == 2{
		if !handlers.PortCheck(os.Args[1]) {
			fmt.Println("enter valid port number")
			return
		}
		port = ":" + os.Args[1]
		fmt.Println("Listening on the port", port)
	}
	handlers.Listiner(port)
}