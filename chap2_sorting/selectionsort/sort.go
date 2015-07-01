// Package selectionsort provides implementation of selection sort
package selectionsort

import (
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
)

func Sort(data sortable.Interface) {
	for i := 0; i < data.Len(); i++ {
		min := i
		for j := i + 1; j < data.Len(); j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}
