package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"net-cat/handlers"
)

/*

Main function to start the TCP chat server
cheak the port from command line arguments
listen on the port and accept incoming connections
handle each client connection in a separate goroutine

*/

func main() {
	var Port string
	// Cheak Command Line Arguments
	if len(os.Args) == 2 {
		Port = os.Args[1]
		num, err := strconv.Atoi(Port)
		if err != nil {
			fmt.Println("Invalide Port Number")
			return
		}
		if num < 1024 {
			fmt.Println("Invalide Port Number chose Port over than 1023")
			return
		}
		// Port is valid
	} else if len(os.Args) == 1 {
		Port = "8989"
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	// Start Listening on the Port
	fmt.Printf("Listening on the port :%s\n", Port)
	ln, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		fmt.Println("We Can Not Conect On this Port")
		return
	}
	// Accept Incoming Connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Errore in Conection Try Again Pleas")
			continue
		}
		// Handle Client Connection
		go handlers.HandleClien(conn)
	}
}
