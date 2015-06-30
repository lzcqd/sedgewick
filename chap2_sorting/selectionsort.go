// Package selectionsort provides implementation of selection sort
package selectionsort

type Sortable interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

func SelectionSort(data Sortable) {
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
