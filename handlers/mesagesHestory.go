package handlers

import "net"

func MesagesHestory(conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()
	for _, val := range allMesages {
		conn.Write([]byte(val + "\n"))
	}
}
