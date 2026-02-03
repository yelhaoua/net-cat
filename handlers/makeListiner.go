package handlers

import (
	"fmt"
	"net"
)

func Listiner(port string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Fialed to make listiner! : ", err)
		return
	}
	for {
		con, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed to connecte !!", err)
			continue
		}
		go HandleClien(con)
	}
}
