package main

import "fmt"

func main() {
	var arr_1 = [5]int{1, 2, 3, 4, 5}
	// 要素(番目)にアクセス
	println(arr_1[0])

	// 要素数を値から推論する
	arr_2 := [...]int{10, 20, 30, 40}
	println(len(arr_2))

	// スライス演算
	fmt.Println((arr_2[0:2]))
}
