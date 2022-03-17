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

	// スライスの初期化
	// 長さと容量を指定して初期化 (各要素の初期値はゼロ値となる)
	var arr_3 []int
	arr_3 = make([]int, 2, 3)
	println(len(arr_3))
	println(arr_3[0])
	println(arr_3[1])
	arr_3 = append(arr_3, 1) // 要素の追加
	println(len(arr_3))
	println(arr_3[2])

	// スライスリテラルで初期化 (要素数の指定は不要かつ自動で配列が作られる)
	var arr_4 = []int{100, 200, 300}
	println(len(arr_4))
	println(cap(arr_4))        // 容量
	arr_4 = append(arr_4, 400) // 要素の追加
	println(cap(arr_4))        // append により容量を超えたら、元の約2倍の容量の配列を確保しなおす

	// マップ
	// リテラルで初期化
	m := map[string]int{"x": 10, "y": 20}
	println(m["x"])

	// キーを指定して入力
	m["z"] = 30
	// 存在確認
	n, ok := m["z"]
	println(n, ok)

	// キーを指定して削除
	delete(m, "z")
	n_, ok := m["z"]
	// 削除結果を確認
	println(n_, ok)
}
