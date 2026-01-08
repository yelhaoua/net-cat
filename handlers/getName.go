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
		conn.Write([]byte("\033[32m[ENTER YOUR NAME]: \033[0m"))
		name, err := bufio.NewReader(conn).ReadString('\n')
		name = strings.Trim(name, "\r\n")
		name = strings.TrimSpace(name)
		if err != nil {
			conn.Close()
		}
		if len(name) > 15 {
			conn.Write([]byte("\033[31m[INVALID USERNAME. IT MUST BE LESS THAN 15 CHARACTERS.]\033[0m\n"))
		} else if name == "" {
			conn.Write([]byte("\033[31m[INVALID USERNAME. IT MUST BE MORE THAN 1 CHARACTER.]\033[0m\n"))
		} else if CheakName(name) {
			conn.Write([]byte("\033[31m[INVALID USERNAME. IT MUST CONTAIN ONLY ENGLISH CHARACTERS.]\033[0m\n"))
		} else if IsExiste(name, conn) {
			conn.Write([]byte("\033[31m[USERNAME IS ALREADY TAKEN.]\033[0m\n"))
		} else {
			mu.Lock()
			user[conn] = strings.TrimSpace(name)
			mu.Unlock()
			return name
		}
	}
}
