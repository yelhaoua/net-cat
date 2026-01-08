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
	if len(user) >= 10 {
		conn.Write([]byte("[ROME IS FULL]"))
		mu.Lock()
		delete(user, conn)
		conn.Close()
		mu.Unlock()
		return
	}
	conn.Write([]byte("\033[33;1m" + baner + "\033[0m"))
	name := GetName(conn)
	fullMsg := fmt.Sprintf("\033[34m%s has joined our chat...\033[0m", name)
	MesagesHestory(conn)
	Send(fullMsg, conn)
	for {
		TM := time.Now().Format("2006-01-02 15:04:05")
		conn.Write([]byte(fmt.Sprintf("[%s][%s]:", TM, name)))
		msg, err := bufio.NewReader(conn).ReadString('\n')
		msg = strings.Trim(msg, "\r\n")
		msg = strings.TrimSpace(msg)
		if err != nil {
			mu.Lock()
			delete(user, conn)
			mu.Unlock()
			fullMsg := fmt.Sprintf("\033[34m%s has left our chat...\033[0m", name)
			Send(fullMsg, conn)
			conn.Close()
			return
		}
		if msg == "" || CheakName(msg) {
			continue
		} else if len(msg) > 1000 {
			conn.Write([]byte("\033[31m[WE CAN’T SEND A MESSAGE OVER 1000 CHARACTERS.]\033[0m\n"))
		} else if strings.HasPrefix(msg, "--NC") {
			arrCommand := strings.Split(msg, " ")
			if len(arrCommand) > 3 {
				conn.Write([]byte("\033[31m[?]	Invalide Command\033[0m\n"))
				continue
			} else if len(arrCommand) == 1 {
				conn.Write([]byte("\033[31m[?]	--NC Reserved Key Word\033[0m\n"))
				conn.Write([]byte("\033[31m[?]	[USE --NC h FOR MORE INFO.]\033[0m\n"))
				continue
			}
			switch arrCommand[1] {
			case "h":
				conn.Write([]byte("\033[33m[IF YOU WANT CHANGE YOU NAME USE --NC ch [NEW USER NAME]]\033[0m\n"))
			case "ch":
				if !CheakName(arrCommand[2]) && !IsExiste(arrCommand[2], conn) {
					mu.Lock()
					oldname := name
					name = arrCommand[2]
					mu.Unlock()
					fullMsg = fmt.Sprintf("\033[34m%s Change His Name TO %s\033[0m", oldname, name)
					Send(fullMsg, conn)
				} else {
					conn.Write([]byte("\033[31m[USERNAME IS ALREADY TAKEN.]\033[0m\n"))
				}
			default:
				conn.Write([]byte(fmt.Sprintf("\033[33m[%s NOT COMMAND IN --NC]\033[0m\n", arrCommand[1])))
			}

		} else {
			fullMsg := fmt.Sprintf("[%s][%s]:[%s]", TM, name, msg)
			Send(fullMsg, conn)
		}
	}
}
