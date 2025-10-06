package main

import "fmt"

const (
	a = 10
	b = 3
)

func main() {
	fmt.Println(a + b) // 13 足し算
	fmt.Println(a - b) // 7　引き算
	fmt.Println(a * b) // 30　掛け算
	fmt.Println(a / b) // 3　割り算
	fmt.Println(a % b) // 1　余り
	fmt.Println(a ^ b) // 1000　累乗
}
