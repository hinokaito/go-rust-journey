package main

import "fmt"

/*
	Goでは、constで宣言された値は、変数に代入されるまで特定の方をもたない、「型付けされない」状態になります。
	カメレオンのような数字と考えるといいです。
*/

const UnTypeNum = 100    // 型を指定していないので「型付けされない定数」
const TypedNum int = 200 // 型を明記したので「int型の定数」

func main() {
	var f float64
	// UnTypeNumは文脈に応じてfloat64として扱われる
	f = UnTypeNum
	fmt.Printf("f: is %T, Value is %v\n", f, f) // f: is float64, Value is 100

	var i int
	// UnTypeNumはここではintとして扱われる
	i = UnTypeNum
	fmt.Printf("i is %T, value is %v\n", i, i) // i is int, value is 100

	// 一方、型が決まってる定数はほかの型には代入できない
	// f = TypedNum  // コンパイルエラー: cannot use TypedNum (type int) as type float64 in assignment
}
