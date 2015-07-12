package quicksort

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
		{sortable.Stringslice([]string{"q", "u", "i", "c", "k", "s", "o", "r", "t"}),
			sortable.Stringslice([]string{"c", "i", "k", "o", "q", "r", "s", "t", "u"})},
	}

	for _, c := range cases {
		Sort(c.in)
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("Sorting result: %v, want: %v", c.in, c.want)
		}
	}
}
