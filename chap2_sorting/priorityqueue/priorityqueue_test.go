package priorityqueue

import (
	"reflect"
	"testing"
)

func TestIntPriorityQueue(t *testing.T) {
	cases := []struct {
		in        []comparable
		maxResult []comparable
		want      PriorityQueue
	}{
		{[]comparable{myint(1), myint(3), myint(4), myint(2), myint(5)},
			[]comparable{myint(5), myint(4), myint(3)},
			(*IntPriorityQueue)(&[]myint{2, 1})},
	}

	for _, c := range cases {
		i := make(IntPriorityQueue, 0)
		for _, v := range c.in {
			(&i).Insert(v)
			t.Log(i)
		}

		if len(i) != len(c.in) {
			t.Errorf("Queue length after insert is not expected, want %v, got %v", len(c.in), len(i))
		}

		for _, r := range c.maxResult {
			t.Log(i)
			g := (&i).DelMax()
			got := reflect.ValueOf(g).Int()
			if r.compare(got) != 0 {
				t.Errorf("DelMax return value unexpected, want: %v, got: %v", r, got)
			}
		}

		if !reflect.DeepEqual(&i, c.want) {
			t.Errorf("Final state of queue unexpected, want: %v, got %v", c.want, &i)
		}
	}
}

func TestStringPriorityQueue(t *testing.T) {
	cases := []struct {
		in        []comparable
		maxResult []comparable
		want      PriorityQueue
	}{
		{[]comparable{mybyte('p'), mybyte('r'), mybyte('i'), mybyte('o'), mybyte('t'), mybyte('y')},
			[]comparable{mybyte('y'), mybyte('t'), mybyte('r')},
			(*StringPriorityQueue)(&[]mybyte{'p', 'o', 'i'})},
	}

	for _, c := range cases {
		s := make(StringPriorityQueue, 0)
		for _, v := range c.in {
			(&s).Insert(v)
			t.Logf("%s", s)
		}

		if len(s) != len(c.in) {
			t.Errorf("Queue length after insert is not expected, want %v, got %v", len(c.in), len(s))
		}

		for _, r := range c.maxResult {
			t.Logf("%s", s)
			g := (&s).DelMax()
			got := mybyte(reflect.ValueOf(g).Uint())
			if r.compare(got) != 0 {
				t.Errorf("DelMax return value unexpected, want: %v, got: %v", r, got)
			}
		}

		if !reflect.DeepEqual(&s, c.want) {
			t.Errorf("Final state of queue unexpected, want: %v, got %v", c.want, &s)
		}
	}
}
