package main

import (
	"fmt"
	"github.com/lzcqd/sedgewick/chap2_sorting/insertionsort"
	"github.com/lzcqd/sedgewick/chap2_sorting/selectionsort"
	"github.com/lzcqd/sedgewick/chap2_sorting/shellsort"
	"github.com/lzcqd/sedgewick/chap2_sorting/sortable"
	"math/rand"
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
		go timesort(s, to_sort, outs[i])
	}

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
				break
			}
		case <-quit:
			fmt.Println("Time out.")
			break
		}
	}
}

func getSortFunc(in string) func(sortable.Interface) {
	switch in {
	case "selectionsort":
		return selectionsort.Sort
	case "insertionsort":
		return insertionsort.Sort
	case "shellsort":
		return shellsort.Sort
	default:
		return nil
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
