package main

import "fmt"

func main() {
	// println("Hello, world!")
	fmt.Print("Hello, ")
	fmt.Print("world with fmt!")
	fmt.Print("a", 111, true)
	// println("Hello, world!")

	// if
	x := 2
	if x == 1 {
		println("x is 1")
	} else if x == 2 {
		println("x is 2")
	} else {
		println("x is neither 1 nor 2.")
	}
}
