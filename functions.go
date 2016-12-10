package main

import (
	"sort"
)

type SortableFloat32 []float32

func (f SortableFloat32) Len() int {
	return len(f)
}

func (f SortableFloat32) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f SortableFloat32) Less(i, j int) bool {
	return f[i] < f[j]
}

func Median(data []float32) float32 {
	sortable := SortableFloat32(data)
	sort.Sort(sortable)
	return float32(sortable[len(sortable)/2])
}

func Average(data []float32) float32 {
	return Sum(data) / float32(len(data))
}

func Sum(data []float32) float32 {
	arraySum := float32(0)
	for _, value := range data {
		arraySum += value
	}
	return arraySum
}
