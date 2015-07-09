package mergesort

import (
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
)

func Sort(data sortable.Interface) {
	var aux sortable.Interface
	aux = data
	for i := 0; i < data.Len(); i++ {
		aux.Add(data.Get(i))
	}

	mergeSort(data, aux, 0, data.Len()-1)
}

func mergeSort(data sortable.Interface, aux sortable.Interface, start int, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(data, aux, start, mid)
	mergeSort(data, aux, mid+1, end)
	merge(data, aux, start, mid, end)
}

func merge(data sortable.Interface, aux sortable.Interface, start int, mid int, end int) {
	for i := start; i <= end; i++ {
		aux.Set(i, data.Get(i))
	}

	i, j, k := start, mid+1, start
	for i <= mid || j <= end {
		if aux.Less(i, j) {
			data.Set(k, aux.Get(i))
			i = i + 1
		} else {
			data.Set(k, aux.Get(j))
			j = j + 1
		}
		k = k + 1
	}

	for i <= mid {
		data.Set(k, aux.Get(i))
		i, k = i+1, k+1
	}

	for j <= end {
		data.Set(k, aux.Get(j))
		j, k = j+1, k+1

	}
}
