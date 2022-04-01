package main

import "fmt"

func main() {
	println("init main")

	print_hello()

	// 戻り値がある関数は変数に代入可能
	msg := hello()
	println(msg)

	add_numbers := add(1, 2)
	println(add_numbers)

	// 多値の受け取り方
	swap_number_a, swap_number_b := swap(1, 2)
	println(swap_number_a)
	println(swap_number_b)
	// 不要な受け取り値は、ブランク変数で受け取り省略可能
	swap_number_x, _ := swap(10, 20)
	_, swap_number_y := swap(10, 20)
	println(swap_number_x)
	println(swap_number_y)

	// 無名関数
	msg_for_clusure := "テストだよ"
	// クロージャー
	func() {
		println(msg_for_clusure)
		// 無名関数を定義した直後に()ですぐ呼び出している
	}()

	// 関数型
	fs := make([]func() string, 2)
	// ↑ string 型を返す関数
	fs[0] = func() string {
		return "foo"
	}
	fs[1] = func() string {
		return "bar"
	}
	for _, f := range fs {
		fmt.Println(f())
	}
}

func print_hello() {
	println("Hello!")
}

func hello() string {
	return "hello"
}

func add(x int, y int) int {
	return x + y
}

// 引数の型をまとめて記述できる
// 複数の戻り値を持たせられる
func swap(x, y int) (int, int) {
	return y, x
}
