package main

import "fmt"

/*
配列は、同じ型のデータを順番に並べたものです。
配列は、「決められた数の、同じ型のデータ」を順序付けて格納するためのものです。
仕切りの数が決まった整理箱のようなものです。

以下の特徴を抑えておく必要があります。

・固定長 一度作成すると長さを変更できない。

・型の一部として長さを持つ。例えば、[3]intと[5]intは別の型です。

・値渡し 関数に配列を渡すと配列全体のデータがコピーされます。大きな配列だと非効率になることがあります。

・インデックス数は0から始まる。
*/

func main() {
	// 配列をゼロ値で宣言
	var arrays [3]int
	fmt.Println(arrays) // [0 0 0]
	// 配列の長さはlen関数で取得できます。
	fmt.Println(len(arrays)) // 3
	// 配列の容量はcap()で取得できます。
	fmt.Println(cap(arrays)) // 3

	// 配列の各要素に値を代入していきます
	arrays[0] = 1
	arrays[1] = 2
	arrays[2] = 3
	fmt.Println(arrays) // [1 2 3]

	// 宣言と同時に初期化もできます
	arrays2 := [3]string{"a", "b", "c"}
	fmt.Println(arrays2) // [a b c]

	// [...]と書くと、初期値の数から長さを自動で推論してくれます。
	arrays3 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("arrays3の型: %T, 値: %v\n", arrays3, arrays3) //
	// arrays3の型: [5]int, 値: [1 2 3 4 5]
}
