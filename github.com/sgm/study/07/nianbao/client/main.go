package main

import (
	"fmt"
	"github.com/sgm/study/07/protocol"
	"net"
)

// socket_stick/client/protocol.go

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		b, _ := protocol.Encode(msg)
		conn.Write([]byte(b))
	}
}
