package main

func init() {
	FunctionRegistry[FunctionTypeAverage] = Average
}

func Average(data []float32) (float32, error) {
	sum, err := Sum(data)
	return sum / float32(len(data)), err
}
