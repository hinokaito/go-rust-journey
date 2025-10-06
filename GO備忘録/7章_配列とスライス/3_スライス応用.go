package main

import "fmt"

func main() {
	// make関数を使って、長さと容量を先に確保する。
	// make(型, 長さ, 容量)
	t := make([]int, 3, 5)
	fmt.Println("t", t)                            // t [0 0 0]
	fmt.Printf("長さ: %d, 容量: %d\n", len(t), cap(t)) // 長さ: 3, 容量: 5

	// スライスから一部を切り出す(スラッシング)
	part := t[1:3]                     // 1から3未満の要素を切り出す。
	fmt.Println("part", part)          // part [0 0]
	fmt.Println("partの容量は", cap(part)) // partの容量は4
}
