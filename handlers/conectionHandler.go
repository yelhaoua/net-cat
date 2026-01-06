package handlers

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handlecon(con net.Conn) {
	defer con.Close()
	reader := bufio.NewReader(con)
	for {
		beyt, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			fmt.Println("Faild to read data")
			return
		}
		fmt.Print(string(beyt))
	}
}

func ConnHandler() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

}
