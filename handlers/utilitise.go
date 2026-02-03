package handlers

import (
	"net"
)

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
		conn.Write([]byte("Errore In Send Message\n"))
		mu.Lock()
		delete(user, conn)
		mu.Unlock()
		return
	}
}

func IsRoomFull() bool {
	mu.Lock()
	defer mu.Unlock()
	return len(user) < 10
}

func IsExiste(name string) bool {
	mu.Lock()
	defer mu.Unlock()
	for _, userName := range user {
		if name == userName {
			return true
		}
	}
	return false
}
