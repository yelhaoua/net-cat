package handlers

import (
	"bufio"
	"net"
	"strings"
)

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

func GetName(conn net.Conn, reader *bufio.Reader) string {
	for {
		fullMsg := "\033[32m[ENTER YOUR NAME]: \033[0m"
		WriteInConnection(conn, fullMsg)
		name, err := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if err != nil {
			conn.Close()
		}
		if len(name) > 15 || name == "" || CheckControlCharacters(name) {
			fullMsg := "\033[31m[INVALID USERNAME. IT MUST BE LESS THAN 15 CHARACTERS. AND OVER THAN 0]\033[0m\n"
			WriteInConnection(conn, fullMsg)
		} else if IsExiste(name) {
			fullMsg := "\033[31m[USERNAME IS ALREADY TAKEN.]\033[0m\n"
			WriteInConnection(conn, fullMsg)
		} else if strings.HasPrefix(name, "--NC") {
			fullMsg := "\033[31m[USERNAME CANNOT BE A RESERVED KEYWORD.]\033[0m\n\033[31m[USE --NC h FOR MORE INFO.]\033[0m\n"
			WriteInConnection(conn, fullMsg)
		} else if !IsRoomFull(){
			fullMsg := "[ROME IS FULL]"
			WriteInConnection(conn, fullMsg)
			conn.Close()
		}else{
			mu.Lock()
			user[conn] = name
			mu.Unlock()
			return name
		}
	}
}
