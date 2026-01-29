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
	TM := time.Now().Format("2006-01-02 15:04:05")
	mu.Lock()
	clien := make(map[net.Conn]string, len(user))
	for c, n := range user {
		clien[c] = n
	}
	allMesages = append(allMesages, msg)
	mu.Unlock()

	// Send the Message to all Clients except the Sender
	for add, name := range clien {
		if conn != add {
			add.Write([]byte(fmt.Sprintf("\n\033[36m%s\033[0m\n", msg)))
			add.Write([]byte(fmt.Sprintf("[%s][%s]: ", TM, name)))
		}
	}
}
