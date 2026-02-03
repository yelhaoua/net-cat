package handlers

import "net"

func HnadleErr(conn net.Conn) {
	mu.Lock()
	delete(user, conn)
	mu.Unlock()

}
