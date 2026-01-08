package handlers

import (
	"bufio"
	"net"
	"strings"
)

func CheakName(name string) bool {
	for _, val := range name {
		if val < ' ' || val > '~' {
			return true
		}
	}
	return false
}

func IsExiste(name string, conn net.Conn) bool {
	for _, userName := range user {
		if name == userName {
			return true
		}
	}
	return false
}

func GetName(conn net.Conn) string {
	for {
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		name, err := bufio.NewReader(conn).ReadString('\n')
		name = strings.Trim(name, "\r\n")
		name = strings.TrimSpace(name)
		if err != nil {
			conn.Close()
		}
		if len(name) > 15 {
			conn.Write([]byte("[INVALID USERNAME. IT MUST BE LESS THAN 15 CHARACTERS.]\n"))
		} else if name == "" {
			conn.Write([]byte("[INVALID USERNAME. IT MUST BE MORE THAN 1 CHARACTER.]\n"))
		} else if CheakName(name) {
			conn.Write([]byte("[INVALID USERNAME. IT MUST CONTAIN ONLY ENGLISH CHARACTERS.]\n"))
		} else if IsExiste(name, conn) {
			conn.Write([]byte("[USERNAME IS ALREADY TAKEN.]\n"))
		} else {
			mu.Lock()
			user[conn] = strings.TrimSpace(name)
			mu.Unlock()
			return name
		}
	}
}
