package mergesort

import (
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
)

func Sort(data sortable.Interface) {
	aux := data.AllocateNew()
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
	i, j := start, mid+1
	for k := start; k <= end; k++ {
		aux.Set(k, data.Get(k))
	}

	for k := start; k <= end; k++ {
		if i > mid {
			data.Set(k, aux.Get(j))
			j = j + 1
		} else if j > end {
			data.Set(k, aux.Get(i))
			i = i + 1
		} else if aux.Less(i, j) {
			data.Set(k, aux.Get(i))
			i = i + 1
		} else {
			data.Set(k, aux.Get(j))
			j = j + 1
		}

	}
}
