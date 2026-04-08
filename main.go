package main

import (
	"fmt"
	format "learing-go/do-format"
	math "learing-go/math"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)

	fmt.Println("==== Go Routines ----")

	processConcurently([]int{1,2,3,4,5})



}


func process(val int) int {
	fmt.Println("processing:", val)
	return val + 10
}

func processConcurently(inVals []int) []int {
	// create cahnnels
	in := make(chan int, 5)
	out := make(chan int, 5)
	n := len(inVals)	
	// launch processing goroutines
	for i := 0; i < n; i ++ {
		go func() {
			for val := range in {
				out <- process(val)
			}
		} ()
	}

	return []int{1,2,3,4}
}



