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

func (i IntPriorityQueue) Insert(key interface{}) {
	v := reflect.ValueOf(key)
	k := myint(int(v.Int()))

	i = append(i, k)
}

func (i IntPriorityQueue) DelMax() interface{} {

}

func (i IntPriorityQueue) toComparable() []comparable {
	c := make([]comparable, len(i))

	for k := 0; k < len(i); k++ {
		c[k] = i[k]
	}

	return c
}

func swim(array []comparable, curr, total int) {

}

func sink(array []comparable, curr, total int) {

}
