package main

/* to run profiling:
go build -cpuprofile cpu.prof -memprofile mem.prof -bench .
go run -cpuprofile cpu.prof -memprofile mem.prof -bench .
*/

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

// go run -cpuprofile=cprof cpu.prof -memprofile=mprof mem.prof -bench .
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	fmt.Println("Merge Sort on an array of positive integers")
	a := []int{4, 3, 1, 2, 5, 3}
	fmt.Println(mergeSort(a))

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

// Merge sort algorithm for an array of positive integers
func mergeSort(a []int) []int {
	// base case: array of length 1
	if len(a) == 1 {
		return a
	} else {
		b := mergeSort(a[0 : len(a)/2])
		c := mergeSort(a[len(a)/2 : len(a)])
		j := 0
		k := 0
		r := make([]int, len(a))
		for i := 0; i < len(a); {

			_b := -1
			_c := -1
			if j < len(b) {
				_b = b[j]
			}
			if k < len(c) {
				_c = c[k]
			}

			if (_b <= _c && _b > -1) || (_b != -1 && _c == -1) {
				r[i] = _b
				j++
				i++
			}

			if (_c <= _b && _c > -1) || (_c != -1 && _b == -1) {
				r[i] = _c
				k++
				i++
			}
		}
		return r
	}
}
