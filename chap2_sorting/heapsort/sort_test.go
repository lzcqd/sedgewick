package heapsort

import (
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	cases := []struct {
		in, want sortable.Interface
	}{
		{sortable.Intslice([]int{8, 3, 5, 7, 10, 1, 4, 2, 9, 6}), sortable.Intslice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})},
		{sortable.Stringslice([]string{"i", "n", "s", "e", "r", "t", "i", "o", "n", "s", "o", "r", "t"}),
			sortable.Stringslice([]string{"e", "i", "i", "n", "n", "o", "o", "r", "r", "s", "s", "t", "t"})},
	}

	for _, c := range cases {
		Sort(c.in)
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("Sorting result: %v, want: %v", c.in, c.want)
		}
	}
}
