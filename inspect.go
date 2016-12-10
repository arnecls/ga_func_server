package main

import (
	"runtime"
)

func init() {
	FunctionRegistry[FunctionTypeCPU] = Cores
	FunctionRegistry[FunctionTypeMemory] = Memory
}

func Cores([]float32) (float32, error) {
	return float32(runtime.NumCPU()), nil
}

func Memory([]float32) (float32, error) {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	return float32(mem.TotalAlloc) / float32(1024.0*1024.0), nil
}
