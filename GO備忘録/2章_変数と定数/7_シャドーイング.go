package main

import "fmt"

/*
	これはバグの原因になりやすい現象です。内側のスコープで、
	外側のスコープにある変数と「同じ名前の変数」を宣言してしまうことを指します。
	多くの静的解析ツールはこれを警告してくれます。
*/

func main() {
	x := 100                  // 外側のスコープのx
	fmt.Println("Before:", x) // Before: 100

	if true {
		// この`x`は外側とは別の、ifブロック内だけの変数
		x := 10                  // シャドーイングが発生
		fmt.Println("Inner:", x) // Inner: 10
	}

	fmt.Println("After:", x) // After: 100 (外側のxは変わっていない！)
}
