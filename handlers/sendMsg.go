package handlers

import (
	"fmt"
	"net"
	"time"
)

/*
send the mesages to all connected clients except the sender
*/

func Send(msg string, conn net.Conn) {
	time := time.Now().Format("2006-01-02 15:04:05")
	
	mu.Lock()
	clien := make(map[net.Conn]string, len(user))
	for c, n := range user {
		clien[c] = n
	}
	mu.Unlock()

	// Send the Message to all Clients except the Sender
	for add, receverName := range clien {
		if conn != add {
			fullMsg := fmt.Sprintf("\n\033[36m%s\033[0m\n[%s][%s]:", msg, time, receverName)
			WriteInConnection(add, fullMsg)
		}
	}
}
