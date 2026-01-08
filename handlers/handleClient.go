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

func HandleClien(conn net.Conn) {
	if len(user) >= 11 {
		conn.Write([]byte("[ROME IS FULL]"))
		mu.Lock()
		delete(user, conn)
		conn.Close()
		mu.Unlock()
		return
	}
	conn.Write([]byte(baner))
	name := GetName(conn)
	fullMsg := fmt.Sprintf("\n%s has joined our chat...", name)
	MesagesHestory(conn)
	Send(fullMsg, conn)
	for {
		TM := time.Now().Format("2006-01-02 15:04:05")
		conn.Write([]byte(fmt.Sprintf("[%s][%s]:", TM, name)))
		msg, err := bufio.NewReader(conn).ReadString('\n')
		msg = strings.Trim(msg, "\r\n")
		if err != nil {
			mu.Lock()
			delete(user, conn)
			mu.Unlock()
			fullMsg := fmt.Sprintf("\n%s has left our chat...", name)
			Send(fullMsg, conn)
			conn.Close()
			return
		}
		if strings.TrimSpace(msg) == "" || CheakName(msg) {
			continue
		} else {
			fullMsg := fmt.Sprintf("[%s][%s]: [%s]", TM, name, msg)
			Send(fullMsg, conn)
		}
	}
}
