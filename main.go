package main

import (
	"fmt"
	"net-cat/handlers"
	"os"
)

/*

Main function to start the TCP chat server
cheak the port from command line arguments
listen on the port and accept incoming connections
handle each client connection in a separate goroutine

*/

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