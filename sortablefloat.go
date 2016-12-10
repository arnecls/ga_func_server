package main

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
