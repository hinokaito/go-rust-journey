package main

import "fmt"

func main() {
	// varを使って、複数の変数を宣言できます。
	var x, y int = 1, 2
	fmt.Println(x, y) // 1 2

	// 変数を宣言するだけでは、ゼロ値が代入されます。
	var a, b int
	fmt.Println(a, b) // 0 0

	// 型の異なる変数を宣言することもできます。
	var str, num = "Hello", 123
	fmt.Println(str, num) // Hello 123
	// str, num := "Hello", 123でもOK

	// ()をつかってまとめることもできます
	var (
		str2  = "Hello"
		num2  = 123
		bool2 = true
		m, n  = 40, "World"
	)
	fmt.Println(str2, num2, bool2, m, n) // Hello 123 true 40 World

}
