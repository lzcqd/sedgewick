package priorityqueue

import (
	"reflect"
)

type PriorityQueue interface {
	Insert(key interface{})
	DelMax() interface{}
}

type comparable interface {
	compare(val interface{}) int
}

type IntPriorityQueue []myint

type myint int

func (i myint) compare(val interface{}) int {
	v := reflect.ValueOf(val)
	c := v.Int()
	if int64(i) < c {
		return -1
	} else if int64(i) == c {
		return 0
	} else {
		return 1
	}
}

func (i *IntPriorityQueue) Insert(key interface{}) {
	v := reflect.ValueOf(key)
	k := myint(int(v.Int()))

	*i = append(*i, k)
	r := (*i).toComparable()
	swim(r, len(*i)-1)

	for j := 0; j < len(*i); j++ {
		(*i)[j] = r[j].(myint)
	}
}

func (i *IntPriorityQueue) DelMax() interface{} {
	r := (*i)[0]
	(*i)[0], (*i)[len(*i)-1] = (*i)[len(*i)-1], (*i)[0]
	*i = (*i)[:len(*i)-1]
	c := (*i).toComparable()

	sink(c, 0, len(*i)-1)

	for j := 0; j < len(*i); j++ {
		(*i)[j] = c[j].(myint)
	}
	return r
}

func (i IntPriorityQueue) toComparable() []comparable {
	c := make([]comparable, len(i))

	for k := 0; k < len(i); k++ {
		c[k] = i[k]
	}

	return c
}

type StringPriorityQueue []mybyte

type mybyte byte

func (b mybyte) compare(val interface{}) int {
	v := mybyte(reflect.ValueOf(val).Uint())
	if b < v {
		return -1
	} else if b == v {
		return 0
	} else {
		return 1
	}
}

func (s *StringPriorityQueue) Insert(val interface{}) {
	v := mybyte(reflect.ValueOf(val).Uint())
	*s = append(*s, v)

	c := (*s).toComparable()
	swim(c, len(c)-1)

	for i := 0; i < len(*s); i++ {
		(*s)[i] = c[i].(mybyte)
	}
}

func (s *StringPriorityQueue) DelMax() interface{} {
	r := (*s)[0]

	(*s)[0], (*s)[len(*s)-1] = (*s)[len(*s)-1], (*s)[0]

	*s = (*s)[:len(*s)-1]

	c := (*s).toComparable()
	sink(c, 0, len(*s)-1)

	for i := 0; i < len(*s); i++ {
		(*s)[i] = c[i].(mybyte)
	}

	return r
}

func (s StringPriorityQueue) toComparable() []comparable {
	c := make([]comparable, len(s))

	for i := 0; i < len(s); i++ {
		c[i] = s[i]
	}
	return c
}

func swim(array []comparable, curr int) {
	for curr > 0 && array[curr].compare(array[(curr-1)/2]) > 0 {
		array[curr], array[(curr-1)/2] = array[(curr-1)/2], array[curr]
		curr = (curr - 1) / 2
	}
}

func sink(array []comparable, curr, total int) {
	for curr*2+1 <= total {
		c := curr*2 + 1
		if c < total && array[c].compare(array[c+1]) < 0 {
			c = c + 1
		}
		if array[curr].compare(array[c]) > 0 {
			break
		}
		array[curr], array[c] = array[c], array[curr]
		curr = c
	}
}
