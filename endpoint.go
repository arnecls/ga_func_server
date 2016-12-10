package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func readHeader(client net.Conn) (*Request, error) {
	rawHeader := make([]byte, 8)
	if _, err := io.ReadFull(client, rawHeader); err != nil {
		return nil, err
	}

	// Convert header
	msg := &Request{}
	reader := bytes.NewReader(rawHeader)
	if err := binary.Read(reader, binary.LittleEndian, &msg.Header); err != nil {
		return nil, err
	}

	fmt.Printf("Header: %#v\n", msg.Header)
	return msg, nil
}

func readData(client net.Conn, msg *Request) error {
	rawData := make([]byte, msg.Header.Count*4)

	if _, err := io.ReadFull(client, rawData); err != nil {
		if err != io.EOF {
			return err
		}
	}

	// Convert data
	reader := bytes.NewReader(rawData)
	msg.Data = make([]float32, msg.Header.Count)
	if err := binary.Read(reader, binary.LittleEndian, &msg.Data); err != nil {
		return err
	}
	fmt.Printf("Data: %#v\n", msg.Data)
	return nil
}

func HandleRequest(client net.Conn) {
	defer client.Close()
	response := make([]byte, 4)
	fmt.Println("Processing request")

	// Read header
	msg, err := readHeader(client)
	if err != nil {
		fmt.Println("Error reading header: ", err)
		return
	}

	err = readData(client, msg)
	if err != nil {
		fmt.Println("Error reading data: ", err)
		return
	}

	result := float32(-1)
	switch msg.Header.Type {
	case FunctionTypeMedian:
		result = Median(msg.Data)
	case FunctionTypeAverage:
		result = Average(msg.Data)
	case FunctionTypeSum:
		result = Sum(msg.Data)
	default:
		fmt.Println("Unknown function type")
	}
	fmt.Println("Result: ", result)

	writer := bytes.NewBuffer(response)
	binary.Write(writer, binary.LittleEndian, result)
	client.Write(response)
}
