package handlers

import (
	"fmt"
	"net"
	"strings"
)

func AplyFlage(msg string, name string, conn net.Conn) {
	arrCommand := strings.Split(msg, " ")
	if len(arrCommand) > 3 {
		conn.Write([]byte("\033[31m[?]	Invalide Command\033[0m\n"))
		return
	} else if len(arrCommand) == 1 {
		conn.Write([]byte("\033[31m[?]	--NC Reserved Key Word\033[0m\n"))
		conn.Write([]byte("\033[31m[?]	[USE --NC h FOR MORE INFO.]\033[0m\n"))
		return
	}
	switch arrCommand[1] {
	case "h":
		conn.Write([]byte("\033[33m[IF YOU WANT CHANGE YOU NAME USE --NC ch [NEW USER NAME]]\033[0m\n"))
		conn.Write([]byte("\033[33m[IF YOU WANT SEE ONLINE USERS USE --NC users]\033[0m\n"))
	case "ch":
		if !CheakName(arrCommand[2]) && !IsExiste(arrCommand[2]) && len(arrCommand[2]) <= 15 && len(arrCommand[2]) > 0 {
			mu.Lock()
			oldname := name
			user[conn] = arrCommand[2]
			name = arrCommand[2]
			mu.Unlock()
			fullMsg := fmt.Sprintf("\033[34m%s Change His Name TO %s\033[0m", oldname, name)
			Send(fullMsg, conn)
		} else {
			conn.Write([]byte("\033[31m[USERNAME IS ALREADY TAKEN.]\033[0m\n"))
		}
	case "users":
		mu.Lock()
		var userList []string
		for _, userName := range user {
			userList = append(userList, userName)
		}
		mu.Unlock()
		conn.Write([]byte("\033[33m[ONLINE USERS]: \n" + strings.Join(userList, "\n") + "\033[0m\n"))
	default:
		conn.Write([]byte(fmt.Sprintf("\033[33m[%s NOT COMMAND IN --NC]\033[0m\n", arrCommand[1])))
	}
}
