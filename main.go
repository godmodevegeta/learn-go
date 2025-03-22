package main

import (
	"fmt"
)

func main(){

	fmt.Println("Hi, welcome to the Experience!\nPlease enter your first name:")
	var fname string
	fmt.Scan(&fname)
	fmt.Printf("Configuring profile for %s...\n", fname)
	beyondmain()
	
}

func beyondmain(){
	fmt.Println("Creating worlds...")
}