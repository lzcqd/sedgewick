package sortable

import (
	"reflect"
)

type Interface interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
	Get(i int) interface{}
	Set(i int, val interface{})
}

type Intslice []int

func (a Intslice) Len() int           { return len(a) }
func (a Intslice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Intslice) Less(i, j int) bool { return a[i] < a[j] }

type Stringslice []string

func (s Stringslice) Len() int           { return len(s) }
func (s Stringslice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Stringslice) Less(i, j int) bool { return s[i] < s[j] }

type Floatslice []float64

func (a Floatslice) Len() int           { return len(a) }
func (a Floatslice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Floatslice) Less(i, j int) bool { return a[i] < a[j] }

type Copyable interface {
	Init(count int)
	Get(i int) interface{}
	Set(i int, val interface{})
}

type CopyableIntslice []int

func (a *CopyableIntslice) Init(count int) {
	*a = make([]int, count)
}
func (a CopyableIntslice) Get(i int) interface{} { return a[i] }
func (a CopyableIntslice) Set(i int, val interface{}) {
	v := reflect.ValueOf(val)
	a[i] = int(v.Int())
}

type CopyableStringslice []string

func (s *CopyableStringslice) Init(count int) {
	*s = make([]string, count)
}
func (s CopyableStringslice) Get(i int) interface{} { return s[i] }
func (s CopyableStringslice) Set(i int, val interface{}) {
	v := reflect.ValueOf(val)
	s[i] = v.String()
}

type CopyableFloatslice []float64

func (a *CopyableFloatslice) Init(count int) {
	*a = make([]float64, count)
}
func (a CopyableFloatslice) Get(i int) interface{} { return a[i] }
func (a CopyableFloatslice) Set(i int, val interface{}) {
	v := reflect.ValueOf(val)
	a[i] = v.Float()
}
