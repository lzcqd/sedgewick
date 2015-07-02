// Package insertionsort provides implementation of insertion sort
package shellsort

import (
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
)

func Sort(data sortable.Interface) {
	h := 1
	for h < data.Len()/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := 0; i < data.Len()-h; i = i + h {
			for j := i + h; j > 0 && data.Less(j, j-h); j = j - h {
				data.Swap(j, j-1)
			}
		}
		h = h / 3
	}
}
