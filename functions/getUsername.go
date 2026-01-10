package functions

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func GetUsername(conn net.Conn, reader *bufio.Reader) {
	for {
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		username, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error in reading username !")
				return
			}
		}
		username = strings.TrimSpace(username)
		if !ValidUsername(username, conn) {
				fmt.Fprint(conn, "[!][INVALID NAME]\n")
		}else{
			mu.Lock()
			Clients[conn] = username
			mu.Unlock()
			break
		}

	}
}
