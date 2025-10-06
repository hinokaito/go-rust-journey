package main

import "fmt"

// 整数型同士のビット演算子

const (
	a = 10
	b = 3
)

func main() {
	fmt.Println(a & b)  // 10 AND
	fmt.Println(a | b)  // 11 OR
	fmt.Println(a ^ b)  // 9 XOR
	fmt.Println(a &^ b) // 2 AND NOT
}
