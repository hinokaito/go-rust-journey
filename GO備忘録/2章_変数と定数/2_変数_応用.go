package main

import "fmt"

func main() {
	// 暗黙的な型宣言。型推論される。
	name := "山田"
	age := 30

	fmt.Println(name, age) // 山田 30

	// 複数の変数を同時に宣言
	class, num := "A", 1
	fmt.Printf("クラス: %s, 番号: %d\n", class, num) // クラス: A, 番号: 1
}
