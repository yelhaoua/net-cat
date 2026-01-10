package functions

import (
	"fmt"
	"net"
)

func SenderMsg(conn net.Conn, msg string) {
	for other, names := range Clients {
		if other != conn {
			_, err := other.Write([]byte("\n" + msg))
			if err != nil {
				fmt.Printf("Write error: %s", err)
			}
			PrinBanner(other, names)
		}
	}
}
