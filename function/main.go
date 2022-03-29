package main

func main() {
	println("init main")

	print_hello()

	// 戻り値がある関数は変数に代入可能
	msg := hello()
	println(msg)
}

func print_hello() {
	println("Hello!")
}

func hello() string {
	return "hello"
}
