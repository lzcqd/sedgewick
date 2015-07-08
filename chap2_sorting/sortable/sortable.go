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

func (a Intslice) Len() int                   { return len(a) }
func (a Intslice) Swap(i, j int)              { a[i], a[j] = a[j], a[i] }
func (a Intslice) Less(i, j int) bool         { return a[i] < a[j] }
func (a Intslice) Get(i int) interface{}      { return nil }
func (a Intslice) Set(i int, val interface{}) {}

type Stringslice []string

func (s Stringslice) Len() int                   { return len(s) }
func (s Stringslice) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s Stringslice) Less(i, j int) bool         { return s[i] < s[j] }
func (s Stringslice) Get(i int) interface{}      { return nil }
func (s Stringslice) Set(i int, val interface{}) {}

type Floatslice []float64

func (a Floatslice) Len() int              { return len(a) }
func (a Floatslice) Swap(i, j int)         { a[i], a[j] = a[j], a[i] }
func (a Floatslice) Less(i, j int) bool    { return a[i] < a[j] }
func (a Floatslice) Get(i int) interface{} { return a[i] }
func (a Floatslice) Set(i int, val interface{}) {
	v := reflect.ValueOf(val)
	a[i] = v.Float()
}
