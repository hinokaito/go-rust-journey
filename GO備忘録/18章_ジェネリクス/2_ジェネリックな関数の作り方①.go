package main

import "fmt"

/*1.シンプルな例: スライスの要素を表示する
どんな型のスライスでも受け取って、その要素を表示する関数を作ってみましょう。
・型制約: any(どんな型でもOK)
・型パラメータ: T
*/

// [T any]で[Tはどんな型でも良い]と宣言
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func main() {
	intSlice := []int{1, 2, 3}
	stringSlice := []string{"a", "b", "c"}

	PrintSlice(intSlice)    // 1 2 3
	PrintSlice(stringSlice) // a b c

	// 関数を呼び出すとき、コンパイラが型を推論してくれるので
	// PrintSlice[int](intSlice) のように型を明記する必要はほとんどない
}

/*
このPrintSlice関数は、intのスライスも、stringのスライスも、その他のどんな型のスライスでも、たった1つの定義で処理できます。
*/
