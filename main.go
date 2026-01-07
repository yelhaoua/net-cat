package main

import (
	"fmt"
	"net"

	"net-cat/handlers"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Chat server running on :8080")

	go handlers.Broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handlers.HandleClient(conn)
	}
}
