package functions

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	WelcomMsg = "Welcome to TCP-Chat!\n         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n __| \".        |\\dS\"qML\n |    `.       | `' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n     `-'       `--'\n"
	Clients   = make(map[net.Conn]string)
	Messages  []string
	mu        sync.Mutex
)

func PrinBanner(conn net.Conn, user string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	_, err := conn.Write([]byte("[" + string(now) + "][" + user + "]: "))
	if err != nil {
		fmt.Printf("Write error: %s", err)
	}
}

func HandelConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	mu.Lock()
	if len(Clients) > 9 {
		fmt.Fprint(conn, "\n[!][THE ROOM IS FULL]")
		mu.Unlock()
		return
	}
	mu.Unlock()
	
	fmt.Fprint(conn, WelcomMsg)
	GetUsername(conn, reader)


	mu.Lock()
	Templat := Clients[conn] + " has join the chat...\n"
	mu.Unlock()

	SenderMsg(conn, Templat)

	for _, msgs := range Messages {
		fmt.Fprint(conn, msgs)
	}

	mu.Lock()
	Messages = append(Messages, Clients[conn]+" has join the chat...\n")
	mu.Unlock()

	for {

		PrinBanner(conn, Clients[conn])
		msg, err := reader.ReadString('\n')

		if err != nil {
			SenderMsg(conn, Clients[conn]+" has left the chat...\n")
			mu.Lock()
			Messages = append(Messages, Clients[conn]+" has left the chat...\n")
			delete(Clients, conn)
			mu.Unlock()
			conn.Close()
			return
		}

		msg = strings.TrimSpace(msg)

		if msg == "" {
			continue
		}

		if strings.HasPrefix(msg, "/nick") {
			ChangName(msg, conn)
			continue
		}

		now := time.Now().Format("2006-01-02 15:04:05")
		Templat := "[" + string(now) + "][" + Clients[conn] + "]:" + "[" + msg + "]" + "\n"
		
		mu.Lock()
		Messages = append(Messages, Templat)
		mu.Unlock()

		SenderMsg(conn, Templat)

	}
}
