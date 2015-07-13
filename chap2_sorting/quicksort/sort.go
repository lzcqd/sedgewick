package quicksort

import (
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
)

func Sort(data sortable.Interface) {
	sortable.RandomShuffle(data)
	quickSort(data, 0, data.Len()-1)
}

func Sort3Way(data sortable.Interface) {
	sortable.RandomShuffle(data)
	quickSort3Way(data, 0, data.Len()-1)
}

func quickSort(data sortable.Interface, start, end int) {
	if start >= end {
		return
	}
	p := partition(data, start, end)
	quickSort(data, start, p-1)
	quickSort(data, p+1, end)
}

func quickSort3Way(data sortable.Interface, start, end int) {
	if start >= end {
		return
	}
	v, lt, i, gt := start, start+1, start+1, end
	for i <= gt {
		if data.Less(i, v) {
			data.Swap(i, lt)
			lt = lt + 1
			i = i + 1
		} else if data.Less(v, i) {
			data.Swap(i, gt)
			gt = gt - 1
		} else {
			i = i + 1
		}
	}
	data.Swap(v, lt-1)
	quickSort3Way(data, start, lt-1)
	quickSort3Way(data, gt+1, end)
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
