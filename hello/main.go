package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	println("Hello, world!")

	fmt.Print("Hello, ")
	fmt.Print("world with fmt!")
	fmt.Println("a", 111, true)

	fmt.Println("quote:", quote.Go())

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

	// for (初期値;継続条件;更新文)
	for i := 0; i <= 10; i = i + 1 {
		fmt.Printf("%d\n", i)
		if i == 8 {
			break
		}
	}
	// for (range を使う繰り返し)
	for _, v := range []string{"a", "b", "c"} {
		fmt.Printf("%s\n", v)
	}
	// ラベルを指定しての break
LOOP:
	for _, w := range []int{1, 2, 3, 4, 5} {
		switch {
		case w == 4:
			fmt.Printf("%d\n", w)
			fmt.Println("Now 4 comes, and break loop.")
			break LOOP
		default:
			fmt.Printf("%d\n", w)
		}
	}
}
