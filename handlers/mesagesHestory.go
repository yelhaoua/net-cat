package handlers

import "net"

/*

log the mesages history to new conected clients

*/

func MesagesHestory(conn net.Conn) {
	mu.Lock()
	msgCopy := make([]string, len(allMesages))
	copy(msgCopy, allMesages)
	mu.Unlock()
	for _, val := range msgCopy {
		WriteInConnection(conn, val+"\n")
	}
}
