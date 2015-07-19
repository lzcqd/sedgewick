package priorityqueue

import (
	"reflect"
)

type PriorityQueue interface {
	Insert(key interface{}) PriorityQueue
	DelMax() (interface{}, PriorityQueue)
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

func (i IntPriorityQueue) Insert(key interface{}) PriorityQueue {
	v := reflect.ValueOf(key)
	k := myint(int(v.Int()))

	i = append(i, k)
	r := i.toComparable()
	swim(r, len(i)-1)

	for j := 0; j < len(i); j++ {
		i[j] = r[j].(myint)
	}
	return i
}

func (i IntPriorityQueue) DelMax() (interface{}, PriorityQueue) {
	r := i[0]
	i[0], i[len(i)-1] = i[len(i)-1], i[0]
	i = append([]myint(nil), i[:len(i)-1]...)
	c := i.toComparable()

	sink(c, 0, len(i)-1)

	for j := 0; j < len(i); j++ {
		i[j] = c[j].(myint)
	}
	return r, i
}

func (i IntPriorityQueue) toComparable() []comparable {
	c := make([]comparable, len(i))

	for k := 0; k < len(i); k++ {
		c[k] = i[k]
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
