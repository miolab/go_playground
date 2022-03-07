package main

import "fmt"

func main() {
	println("Hello, world!")

	fmt.Print("Hello, ")
	fmt.Print("world with fmt!")
	fmt.Println("a", 111, true)

	// 書式指定（\n →改行）
	fmt.Printf("Good morning.\n")
	// 書式指定（%d →整数, %s →文字列）
	fmt.Printf("%d...%s\n", 100, "偶数")

	// if
	x := 2
	if x == 1 {
		println("x is 1.")
	} else if x == 2 {
		println("x is 2.")
	} else {
		println("x is neither 1 nor 2.")
	}

	// switch
	a := 0
	switch a {
	case 0:
		fmt.Println("a is 0.")
	case 1, 2:
		fmt.Println("a is 1 or 2.")
	default:
		fmt.Println("default.")
	}
}
