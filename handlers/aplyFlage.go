package handlers

import (
	"fmt"
	"net"
	"strings"
)

/*
AplyFlage apliy the --NC flage commands
--NC h : show help about --NC commands
--NC ch [NEW USER NAME] : change the user name to NEW USER NAME
--NC users : show the list of online users
*/

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

	// Handle Commands

	switch arrCommand[1] {

	// Help Command

	case "h":
		fullMsg := "\033[33m[IF YOU WANT CHANGE YOU NAME USE --NC ch [NEW USER NAME]]\033[0m\n\033[33m[IF YOU WANT SEE ONLINE USERS USE --NC users]\033[0m\n"
		WriteInConnection(conn, fullMsg)

	// Change Name Command

	case "ch":
		if len(arrCommand) != 3 {
			fullMsg := "\033[31m[?]	Invalide Command USE --NC ch [NAME]\033[0m\n"
			WriteInConnection(conn, fullMsg)
			return
		}
		if !IsExiste(arrCommand[2]) && len(arrCommand[2]) <= 15 && len(arrCommand[2]) > 0 {
			mu.Lock()
			oldname := name
			user[conn] = arrCommand[2]
			mu.Unlock()
			name = arrCommand[2]
			fullMsg := fmt.Sprintf("\033[34m%s Change His Name TO %s\033[0m", oldname, name)
			Send(fullMsg, conn)
			fullMsg = fmt.Sprintf("\033[32m[YOUR NAME CHANGED TO %s SUCCESSFULY]\033[0m\n", name)
			WriteInConnection(conn, fullMsg)
		} else {
			fullMsg := "\033[31m[USERNAME IS ALREADY TAKEN.]\033[0m\n"
			WriteInConnection(conn, fullMsg)
		}

	// Show Online Users Command

	case "users":
		mu.Lock()
		var userList []string
		for _, userName := range user {
			userList = append(userList, userName)
		}
		mu.Unlock()
		fullMsg := "\033[33m[ONLINE USERS]: \n" + strings.Join(userList, "\n") + "\033[0m\n"
		WriteInConnection(conn, fullMsg)

	// Invalid Command
	default:
		fullMsg := fmt.Sprintf("\033[33m[%s NOT COMMAND IN --NC]\033[0m\n", arrCommand[1])
		WriteInConnection(conn, fullMsg)
	}
}
