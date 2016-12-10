package main

import (
	"fmt"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":5880")
	if err != nil {
		panic(fmt.Errorf("Listen failed %s", err))
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err)
		}
		go HandleRequest(conn)
	}
}
