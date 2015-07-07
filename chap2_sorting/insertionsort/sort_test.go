package insertionsort

import (
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
	"reflect"
	"testing"
)

type intslice []int

func (a intslice) Len() int                   { return len(a) }
func (a intslice) Swap(i, j int)              { a[i], a[j] = a[j], a[i] }
func (a intslice) Less(i, j int) bool         { return a[i] < a[j] }
func (a intslice) Get(i int) interface{}      { return nil }
func (a intslice) Set(i int, val interface{}) {}

type stringslice []string

func (s stringslice) Len() int                   { return len(s) }
func (s stringslice) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s stringslice) Less(i, j int) bool         { return s[i] < s[j] }
func (s stringslice) Get(i int) interface{}      { return nil }
func (s stringslice) Set(i int, val interface{}) {}

func TestSort(t *testing.T) {
	cases := []struct {
		in, want sortable.Interface
	}{
		{intslice([]int{8, 3, 5, 7, 10, 1, 4, 2, 9, 6}), intslice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})},
		{stringslice([]string{"i", "n", "s", "e", "r", "t", "i", "o", "n", "s", "o", "r", "t"}),
			stringslice([]string{"e", "i", "i", "n", "n", "o", "o", "r", "r", "s", "s", "t", "t"})},
	}

	for _, c := range cases {
		Sort(c.in)
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("Sorting result: %v, want: %v", c.in, c.want)
		}
	}
}
