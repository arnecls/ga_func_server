package main

type FunctionType uint32

const (
	FunctionTypeMedian  = FunctionType(iota)
	FunctionTypeAverage = FunctionType(iota)
	FunctionTypeSum     = FunctionType(iota)
	FunctionTypeCPU     = FunctionType(iota)
	FunctionTypeMemory  = FunctionType(iota)
)

type CallableFunction func([]float32) (float32, error)
type CallableFunctionMap map[FunctionType]CallableFunction

var FunctionRegistry = make(CallableFunctionMap)
