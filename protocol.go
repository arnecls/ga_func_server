package main

type FunctionType uint32

const (
	FunctionTypeMedian  = FunctionType(iota)
	FunctionTypeAverage = FunctionType(iota)
	FunctionTypeSum     = FunctionType(iota)
)

type RequestHeader struct {
	Type  FunctionType
	Count uint32
}

type Request struct {
	Header RequestHeader
	Data   []float32
}

type Response struct {
	Data float32
}
