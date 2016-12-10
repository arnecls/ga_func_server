package main

func init() {
	FunctionRegistry[FunctionTypeSum] = Sum
}

func Sum(data []float32) (float32, error) {
	arraySum := float32(0)
	for _, value := range data {
		arraySum += value
	}
	return arraySum, nil
}
