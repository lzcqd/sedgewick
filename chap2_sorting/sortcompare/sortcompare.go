package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/lzcqd/sedgewick/chap2_sorting/insertionsort"
	"github.com/lzcqd/sedgewick/chap2_sorting/selectionsort"
	"github.com/lzcqd/sedgewick/chap2_sorting/shellsort"
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)

type floatslice []float64

func (a floatslice) Len() int              { return len(a) }
func (a floatslice) Swap(i, j int)         { a[i], a[j] = a[j], a[i] }
func (a floatslice) Less(i, j int) bool    { return a[i] < a[j] }
func (a floatslice) Get(i int) interface{} { return a[i] }
func (a floatslice) Set(i int, val interface{}) {
	v := reflect.ValueOf(val)
	a[i] = v.Float()
}

func timesort(sort func(sortable.Interface), to_sort []floatslice, out chan string) {
	defer close(out)

	start := time.Now()
	for _, s := range to_sort {
		sort(s)
	}

	duration := time.Since(start)
	out <- fmt.Sprintf("%s completed in %v", runtime.FuncForPC(reflect.ValueOf(sort).Pointer()).Name(), duration)
}

func merge(cs ...chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	output := func(c <-chan string) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func startsorts(sorts []func(sortable.Interface), to_sort []floatslice, timeout int) {
	outs := make([]chan string, len(sorts))
	for i := range outs {
		outs[i] = make(chan string)
	}

	out := merge(outs...)
	done := make(chan bool)
	go manageOutput(out, timeout, done)

	for i, s := range sorts {
		new_sort := make([]floatslice, len(to_sort))
		for j := range new_sort {
			var fs []float64
			fs = append(fs, to_sort[j]...)
			new_sort[j] = fs
		}
		go timesort(s, new_sort, outs[i])
	}
	fmt.Println("sorting...")
	<-done
}

func manageOutput(out <-chan string, timeout int, done chan bool) {
	defer func() {
		done <- true
	}()

	quit := time.Tick(time.Duration(timeout) * time.Second)

	for {
		select {
		case s, ok := <-out:
			fmt.Println(s)
			if !ok {
				return
			}
		case <-quit:
			fmt.Println("Time out.")
			return
		}
	}
}

func getSortFunc(in string) (func(sortable.Interface), error) {
	switch in {
	case "selectionsort":
		return selectionsort.Sort, nil
	case "insertionsort":
		return insertionsort.Sort, nil
	case "shellsort":
		return shellsort.Sort, nil
	default:
		return nil, errors.New("fail to parse sort function")
	}
}

func generateSortArray(arrayCount, elementCount int) []floatslice {
	ret := make([]floatslice, arrayCount)
	for i := range ret {
		array := make([]float64, elementCount)
		rand.Seed(int64(time.Now().Unix()))
		for j := range array {
			array[j] = rand.Float64()
		}
		ret[i] = array
	}
	return ret
}

func main() {
	sortFuncs := flag.String("sorts", "", "Sorting functions to compare, comma separated")
	arrayCount := flag.Int("array", 1, "Number of arrays to sort, default 1")
	elementCount := flag.Int("element", 1000, "Number of random entries for each array, default 1000")
	timeout := flag.Int("timeout", 60, "Maximum time to run in seconds, default 60s")
	flag.Parse()

	funcs := strings.Split(*sortFuncs, ",")
	for i, f := range funcs {
		funcs[i] = strings.TrimSpace(f)
	}

	var sorts []func(sortable.Interface)

	for _, f := range funcs {
		s, err := getSortFunc(f)
		if err != nil {
			panic(fmt.Sprintf("Not recognised sort function: %s\n", s))
		}
		sorts = append(sorts, s)
	}

	sortArray := generateSortArray(*arrayCount, *elementCount)
	startsorts(sorts, sortArray, *timeout)
}
