package structs

import "net"

type Client struct {
	Name string
	Conn net.Conn
}
