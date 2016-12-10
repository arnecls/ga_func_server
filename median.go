package main

import (
	"sort"
)

func init() {
	FunctionRegistry[FunctionTypeMedian] = Median
}

func Median(data []float32) (float32, error) {
	sortable := SortableFloat32(data)
	sort.Sort(sortable)
	return float32(sortable[len(sortable)/2]), nil
}
