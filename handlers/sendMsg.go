package handlers

import (
	"fmt"
	"net"
	"time"
)

func Send(msg string, conn net.Conn) {
	TM := time.Now().Format("2006-01-02 15:04:05")

	for add, name := range user {
		if conn != add {
			add.Write([]byte(fmt.Sprintf("\n\033[36m%s\033[0m\n", msg)))
			add.Write([]byte(fmt.Sprintf("[%s][%s]: ", TM, name)))
		}
	}

	mu.Lock()
	allMesages = append(allMesages, msg)
	mu.Unlock()
}
