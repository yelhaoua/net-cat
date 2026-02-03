package handlers

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	baner      = "Welcome to TCP-Chat!\n         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n __| \".        |\\dS\"qML\n |    `.       | `' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n     `-'       `--'\n"
	user       = make(map[net.Conn]string)
	mu         = sync.Mutex{}
	allMesages []string
)

/*

HandleClien handle the client connection
and redirect the messages to Send function

*/

func HandleClien(conn net.Conn) {
	reader := bufio.NewReader(conn)
	conn.Write([]byte("\033[33;1m" + baner + "\033[0m"))
	name := GetName(conn, reader)
	if name == "[ROME IS FULL]" {
		return
	}
	fullMsg := fmt.Sprintf("\033[34m%s has joined our chat...\033[0m", name)
	MesagesHestory(conn)
	Send(fullMsg, conn)

	for {
		TM := time.Now().Format(time.DateTime)
		_, connErr := conn.Write([]byte(fmt.Sprintf("[%s][%s]:", TM, user[conn])))
		msg, err := reader.ReadString('\n')
		if connErr != nil || err != nil {
			HnadleErr(conn)
			fullMsg := fmt.Sprintf("\033[34m%s has left our chat...\033[0m", name)
			Send(fullMsg, conn)
			conn.Close()
			return
		}
		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue
		} else if len(msg) > 1000 {
			_, errconn := conn.Write([]byte("\033[31m[WE CAN’T SEND A MESSAGE OVER 1000 CHARACTERS.]\033[0m\n"))
			if errconn != nil {
				HnadleErr(conn)
				return
			}
		} else if strings.HasPrefix(msg, "--NC") {
			AplyFlage(msg, name, conn)
		} else {
			fullMsg := fmt.Sprintf("[%s][%s]:%s", TM, user[conn], msg)
			Send(fullMsg, conn)
		}
	}
}
