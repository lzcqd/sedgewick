package quicksort

import (
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
)

func Sort(data sortable.Interface) {
	data.RandomShuffle()
	quickSort(data, 0, data.Len()-1)
}

func quickSort(data sortable.Interface, start, end int) {
	if start >= end {
		return
	}
	p := partition(data, start, end)
	quickSort(data, start, p-1)
	quickSort(data, p+1, end)
}

func partition(data sortable.Interface, start, end int) int {
	p, i, j := start, start, end+1
	for true {
		i = i + 1
		for data.Less(i, p) {
			if i >= end {
				break
			}
			i = i + 1
		}
		j = j - 1
		for data.Less(p, j) {
			if j <= start {
				break
			}
			j = j - 1
		}

		if i >= j {
			break
		}
		data.Swap(i, j)
	}
	data.Swap(p, j)
	return j
}
