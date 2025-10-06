package main

import "fmt"

func main() {
	// int: 符号付整数(プラス,マイナス,ゼロ)
	var intValue int
	fmt.Println(intValue) // 0

	// uint: 符号なし整数(プラス,ゼロ)
	var uintValue uint
	fmt.Println(uintValue) // 0

	// float: 浮動小数点数
	var floatValue32 float32  // 32ビット
	var floatValue64 float64  // 64ビット 基本的にfloat64を使う
	fmt.Println(floatValue32) // 0
	fmt.Println(floatValue64) // 0

	// 文字列型
	var stringValue string
	fmt.Println(stringValue) // ""

	// rune: 文字(Unicodeコードポイント)
	var runeValue rune
	fmt.Println(runeValue) // 0

	// bool: 真偽値
	var boolValue bool
	fmt.Println(boolValue) // false
}
