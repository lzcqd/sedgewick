// Package main provides ...
package selectionsort

import (
	"testing"
)

type ascending []int

func (a ascending) Len() int           { return len(a) }
func (a ascending) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ascending) Less(i, j int) bool { return a[i] < a[j] }

func TestSelectionSort(t *testing.T) {
	cases := []struct {
		in, want ascending
	}{
		{ascending([]int{8, 3, 5, 7, 10, 1, 4, 2, 9, 6}), ascending([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})},
	}

	for _, c := range cases {
		SelectionSort(c.in)
		if !equal(c.in, c.want) {
			t.Errorf("Sorting result: %d, want: %d", c.in, c.want)
		}
	}
}

func equal(a, b ascending) bool {
	if a.Len() != b.Len() {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
