package main

import "fmt"

/*
interface{}は、メソッドを1つも定義していない特殊なインターフェースです。
メソッドが1つもないので、Goの全ての型がこのインターフェースを暗黙的に満たします。

つまり、interface{}型の変数には、どんな型の値でも代入できます。
これらは非常に便利ですが、型が何であるか分からなくなるため、乱用は避けるべきです。
fmt.Printlnの引数がこのinterface{}です。
*/

func showType(x interface{}) {
	fmt.Printf("値: %v, 型: %T\n", x, x)
}

func main() {
	showType(100)     // 値: 100, 型: int
	showType("Hello") // 値: Hello, 型: string
	showType(true)    // 値: true, 型: bool
}
