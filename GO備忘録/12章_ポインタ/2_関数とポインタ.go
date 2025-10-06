// ポインタの最も一般的な使い方が、関数で引数の値を変更するケースです。
// Go言語では、ポインタを経由しないと、関数内で変数の値を変更することができません。
package main

import "fmt"

// 値渡し(コピー)なので、元の値は変わらない
func doubleValue(n int) {
	n = n * 2
}

// ポインタ渡しなので、元の値が変わる
func doublePointer(n *int) {
	*n = *n * 2 // ポインタが指す先の値を2倍にする
}

func main() {
	// 値渡しの場合
	val := 10
	doubleValue(val)
	fmt.Printf("値渡しの場合、元の値は変わらない: %d\n", val) // 値渡し後の値: 10

	// ポインタ渡しの場合
	ptrVal := 10
	doublePointer(&ptrVal)
	fmt.Printf("ポインタ渡しの場合、元の値は変わる: %d\n", ptrVal) // ポインタ渡し後の値: 20
}
