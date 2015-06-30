// Package main provides ...
package selectionsort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{8, 3, 5, 7, 10, 1, 4, 2, 9, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, c := range cases {
		result := SelectionSort(c.in)
		if !equal(result, c.want) {
			t.Errorf("Sorting result: %d, want: %d", result, c.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
