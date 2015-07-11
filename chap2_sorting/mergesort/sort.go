package mergesort

import (
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
	"math"
)

func Sort(data sortable.Interface) {
	aux := data.AllocateNew()
	mergeSort(data, aux, 0, data.Len()-1)
}

func SortBU(data sortable.Interface) {
	aux := data.AllocateNew()
	mergeSortBU(data, aux, 0, data.Len()-1)
}

func mergeSort(data, aux sortable.Interface, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(data, aux, start, mid)
	mergeSort(data, aux, mid+1, end)
	merge(data, aux, start, mid, end)
}

func mergeSortBU(data, aux sortable.Interface, start, end int) {
	for step := 1; step <= end; step = step * 2 {
		for i := 0; i <= end-step; i = i + step*2 {
			merge(data, aux, i, i+step-1, int(math.Min(float64(i+step*2-1), float64(end))))
		}
	}
}

func merge(data, aux sortable.Interface, start, mid, end int) {
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
