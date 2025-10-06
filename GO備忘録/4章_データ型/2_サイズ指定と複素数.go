package main

import "fmt"

func main() {
	// メモリ使用量を厳密に管理したい場合や、特定のデータ形式を扱う際には、
	// intやuintのサイズを明示した型を使います。

	var intValue8 int8
	var intValue16 int16
	var intValue32 int32
	var intValue64 int64

	// 異なるサイズ同士で計算したい場合、変換が必要になります。
	// intValue8 += int8(intValue16)など

	var uintValue8 uint8
	var uintValue16 uint16
	var uintValue32 uint32
	var uintValue64 uint64

	// byte: uint8の別名。1バイトのデータを扱う
	var byteValue byte

	// rune: int32の別名。1文字のデータを扱う。日本語や😀などを、1文字として扱えるようになります。
	var runeValue rune

	/*
		もしruneを使わない場合、日本語や😀などを1文字として扱えないので、
		for文などで文字列を分割しようとすると文字化けすることがあります。
	*/

	// complex: 複素数を扱う
	var complexValue64 complex64
	var complexValue128 complex128

	fmt.Println(intValue8)
	fmt.Println(intValue16)
	fmt.Println(intValue32)
	fmt.Println(intValue64)
	fmt.Println(uintValue8)
	fmt.Println(uintValue16)
	fmt.Println(uintValue32)
	fmt.Println(uintValue64)
	fmt.Println(byteValue)
	fmt.Println(runeValue)
	fmt.Println(complexValue64)
	fmt.Println(complexValue128)
}
