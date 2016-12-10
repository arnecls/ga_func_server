package main

import (
	"bytes"
	"encoding/binary"
	"io"
)

type RequestHeader struct {
	Type  FunctionType
	Count uint32
}

type Request struct {
	RequestHeader
	Data []float32
}

type Response struct {
	Data     float32
	Errocode uint32
}

func ReadRequest(in io.Reader) (*Request, error) {
	req, err := readHeader(in)
	if err != nil {
		return nil, err
	}

	err = readData(in, req)
	if err != nil {
		return nil, err
	}

	return req, err
}

func WriteResponse(out io.Writer, rsp Response) error {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.LittleEndian, rsp); err != nil {
		return err
	}

	_, err := out.Write(buffer.Bytes())
	return err
}

func readHeader(stream io.Reader) (*Request, error) {
	rawHeader := make([]byte, binary.Size(RequestHeader{}))
	if _, err := io.ReadFull(stream, rawHeader); err != nil {
		return nil, err
	}

	req := &Request{}
	reader := bytes.NewReader(rawHeader)
	if err := binary.Read(reader, binary.LittleEndian, &req.RequestHeader); err != nil {
		return nil, err
	}

	return req, nil
}

func readData(stream io.Reader, req *Request) error {
	rawData := make([]byte, req.Count*4)
	if _, err := io.ReadFull(stream, rawData); err != nil {
		if err != io.EOF {
			return err
		}
	}

	reader := bytes.NewReader(rawData)
	req.Data = make([]float32, req.Count)
	if err := binary.Read(reader, binary.LittleEndian, &req.Data); err != nil {
		return err
	}

	return nil
}
