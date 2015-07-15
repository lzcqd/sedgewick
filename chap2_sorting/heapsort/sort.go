package heapsort

import (
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
)

func Sort(data sortable.Interface) {
	n := data.Len() - 1
	for i := n / 2; i >= 0; i-- {
		sink(data, i, n)
	}

	for n > 0 {
		data.Swap(0, n)
		n = n - 1
		sink(data, 0, n)
	}
}

func sink(data sortable.Interface, curr, total int) {
	for curr*2+1 <= total {
		c := curr*2 + 1
		if c+1 <= total && data.Less(c, c+1) {
			c = c + 1
		}
		if !data.Less(curr, c) {
			break
		}
		data.Swap(curr, c)
		curr = c
	}
}
