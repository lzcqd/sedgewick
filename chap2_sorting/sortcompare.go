package main

import (
	"fmt"
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
	"reflect"
	"runtime"
	"sync"
	"time"
)

type floatslice []float64

func (a floatslice) Len() int           { return len(a) }
func (a floatslice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a floatslice) Less(i, j int) bool { return a[i] < a[j] }

func timesort(sort func(sortable.Interface), to_sort []floatslice, out chan string) {
	defer close(out)

	start := time.Now()
	for _, s := range to_sort {
		sort(s)
	}

	duration := time.Since(start)
	out <- fmt.Sprintf("%s completed in %v milliseconds", runtime.FuncForPC(reflect.ValueOf(sort).Pointer()).Name(), duration/time.Millisecond)
}

func merge(cs ...<-chan string) <-chan string {
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

//func startsorts
