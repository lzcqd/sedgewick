// Package insertionsort provides implementation of insertion sort
package insertionsort

import (
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
)

func Sort(data sortable.Interface) {
	for i := 0; i < data.Len()-1; i++ {
		for j := i + 1; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
