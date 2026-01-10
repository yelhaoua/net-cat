package functions

import (
	"fmt"
	"net"
	"strings"
)

func ChangName(msg string, conn net.Conn) {
	Splited := strings.Split(msg, " ")
	if len(Splited) != 2 {
		fmt.Fprint(conn, "[!][INVALID NAME : /nick <newname>]\n")
		return
	}
	mu.Lock()
	text := Clients[conn] + " has change his name to " + Splited[1]
	Messages = append(Messages, text+"\n")
	Clients[conn] = Splited[1]
	mu.Unlock()

	SenderMsg(conn, text)
}
