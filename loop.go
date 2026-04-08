package main

import (
	"fmt"
)

func main1() {
	x := []int{1,2,3,4,5}
	for _, v := range x {
		fmt.Printf("%p\n", &v)
	}
}


