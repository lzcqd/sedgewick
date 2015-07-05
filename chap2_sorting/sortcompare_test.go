package main

import (
	"testing"
)

func TestGenerateSortArray(t *testing.T) {
	cases := []struct {
		arrayCount, elementCount int
	}{
		{1, 10},
		{10, 1000},
	}

	for _, c := range cases {
		got := generateSortArray(c.arrayCount, c.elementCount)
		if len(got) != c.arrayCount {
			t.Errorf("Array length not expected. Got %d, want %d", len(got), c.arrayCount)
		}
		for _, s := range got {
			for i := range s {
				t.Logf("%v\n", s[i])
			}
		}
	}
}
