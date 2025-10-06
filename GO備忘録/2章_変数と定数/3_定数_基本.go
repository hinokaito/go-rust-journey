package main

import "fmt"

// // 定数は、再代入不可能な読み取り専用の箱です。

func main() {
	// 基準(プログラム中で変わることが無い)
	const MaximumSpeed = 100

	fmt.Printf("最高速度は%dkm/hです。\n", MaximumSpeed) // 最高速度は100km/hです。

	// 再代入しようとするとエラーになる
	// MaximumSpeed = 120
	//annot assign to MaximumSpeed (neither addressable nor a map index expression)

	/*
		なお、「:=」は変数宣言に使うので、定数には使えない。
		さらに、宣言と同時に値を設定する必要がある。
	*/
}
