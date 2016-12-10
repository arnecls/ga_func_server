package main

import (
	"fmt"
	"net"
	"time"
)

func HandleRequest(client net.Conn) {
	defer func() {
		fmt.Println(client.RemoteAddr(), "Closed")
		client.Close()
	}()
	rsp := Response{
		Data:     0,
		Errocode: 0,
	}

	fmt.Println(client.RemoteAddr(), "Reading request...")
	req, err := receiveRequest(client)
	if err != nil {
		fmt.Println(client.RemoteAddr(), "Error reading request: ", err)
		rsp.Errocode = 1
	} else {
		fmt.Println(client.RemoteAddr(), "req:", req.Type, req.Count, req.Data)
		fmt.Println(client.RemoteAddr(), "Processing...")
		rsp.Data, err = processRequest(req)
		if err != nil {
			fmt.Println(client.RemoteAddr(), "Error processing request: ", err)
			rsp.Errocode = 2
		}
	}

	fmt.Println(client.RemoteAddr(), "Sending result...")
	fmt.Println(client.RemoteAddr(), "rsp:", rsp.Data, rsp.Errocode)
	if err := sendResponse(client, rsp); err != nil {
		fmt.Println(client.RemoteAddr(), "Error sending response: ", err)
		return
	}
}

func receiveRequest(client net.Conn) (*Request, error) {
	client.SetReadDeadline(time.Now().Add(5 * time.Second))
	return ReadRequest(client)
}

func processRequest(req *Request) (float32, error) {
	callable, isKnown := FunctionRegistry[req.Type]
	if !isKnown {
		return 0.0, fmt.Errorf("Unknown function type %X", int(req.Type))
	}
	return callable(req.Data)
}

func sendResponse(client net.Conn, result Response) error {
	client.SetWriteDeadline(time.Now().Add(5 * time.Second))
	return WriteResponse(client, result)
}
