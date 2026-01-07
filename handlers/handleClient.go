package handlers

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"net-cat/structs"
)

var (
	Clients   = make(map[net.Conn]*structs.Client)
	Broadcast = make(chan string)
)

func HandleClient(conn net.Conn) {
	defer conn.Close()

	client := GetName(conn)
	if client == nil {
		return
	}

	joinMsg := fmt.Sprintf("\n[%s] %s joined the chat",
		time.Now().Format("2006-01-02 15:04:05"),
		client.Name)
	Broadcast <- joinMsg

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		msg := fmt.Sprintf("\n[%s][%s]: %s",
			time.Now().Format("2006-01-02 15:04:05"),
			client.Name,
			text)

		Broadcast <- msg
	}

	delete(Clients, conn)
	leaveMsg := fmt.Sprintf("[%s] %s left the chat",
		time.Now().Format("2006-01-02 15:04:05"),
		client.Name)
	Broadcast <- leaveMsg
}

func GetName(conn net.Conn) *structs.Client {
	reader := bufio.NewReader(conn)

	for {
		conn.Write([]byte("[Enter Your Name]: "))
		name, err := reader.ReadString('\n')
		if err != nil {
			return nil
		}
		name = strings.TrimSpace(name)

		if name != "" {
			client := &structs.Client{
				Name: name,
				Conn: conn,
			}
			Clients[conn] = client
			return client
		}
	}
}

func Broadcaster() {
	for msg := range Broadcast {
		for _, client := range Clients {
			fmt.Fprintln(client.Conn, msg)
			fmt.Fprintf(client.Conn, "[%s][%s]: ",
				time.Now().Format("2006-01-02 15:04:05"),
				client.Name,
			)
		}
	}
}
