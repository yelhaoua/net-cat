package handlers

import "net"

func MesagesHestory(conn net.Conn) {
	mu.Lock()
	msgCopy := make([]string, len(allMesages))
	copy(msgCopy, allMesages)
	mu.Unlock()
	for _, val := range msgCopy {
		conn.Write([]byte(val + "\n"))
	}
}
