package handlers

import "net"

func CheckControlCharacters(message string) bool {
	for _, val := range message {
		if val <= 31 {
			return true
		}
	}
	return false
}

func WriteInConnection(conn net.Conn, messege string) {
	_, err := conn.Write([]byte(messege))
	if err != nil {
		mu.Lock()
		delete(user, conn)
		mu.Unlock()
		return
	}
}
