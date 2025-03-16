package main

import (
	"fmt"
	"reflect"
)

func main(){

	fmt.Println("Hi, there!")

	var a = "hello world :)"

	fmt.Println(reflect.TypeOf(a))

	b := 90

	b = 88
	
	fmt.Println(b)
}