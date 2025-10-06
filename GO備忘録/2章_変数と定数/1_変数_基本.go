package main

import "fmt"

// 変数は、ご存じの通り再代入可能な箱です。
// Goは静的型付け言語なので、型を指定した後別の型のものを代入することはできません。
// なお、Goのルールとして、宣言した変数は全て使わなければならないというルールがあります。

func main() {
	// var 変数名 型 = 初期値
	var name string = "山田"
	fmt.Println(name) // 山田

	// 型だけ明記する(ゼロ値)
	var age int
	var bloodType string
	fmt.Println(age)       // 0
	fmt.Println(bloodType) // ゼロ値(空文字列)

	// var 変数名 = 初期値(型推論)
	var city = "東京"   // 東京は文字列なのでstring型だと推論される
	fmt.Println(city) // 東京
}
