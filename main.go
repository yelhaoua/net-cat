package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"net-cat/handlers"
)

func main() {
	var Port string

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
	} else if len(os.Args) == 1 {
		Port = "8989"
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	
	fmt.Printf("Listening on the port :%s", Port)
	ln, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		fmt.Println("We Can Not Conect On this Port")
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Errore in Conection Try Again Pleas")
			continue
		}
		go handlers.HandleClien(conn)
	}
}
