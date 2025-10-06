package main

import "fmt"

/*2.型制約を使った例: 数値を合計する
次に、intやfloat64など、数値のスライスを受け取って合計を計算する関数を考えます。

ただのanyでは足し算(+)ができるか分からないため、数値型に限定する制約が必要です。
Goでは、constraintsという便利な標準パッケージが用意されています。
*/

// Go 1.18以降では、標準ライブラリの制約を使用
// Numberという制約を独自に定義する
// 数値型のみを受け入れる制約を定義
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// SumNumbers は、制約として定義した Number を満たす型のみを受け取る
func SumNumbers[T Number](nums []T) T {
	var total T // T型のゼロ値で初期化 (intなら0, float64なら0.0)
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	intSlice := []int{1, 2, 3, 4}
	floatSlice := []float64{1.1, 2.2, 3.3}

	fmt.Println("Sum of ints:", SumNumbers(intSlice))     // Sum of ints: 10
	fmt.Println("Sum of floats:", SumNumbers(floatSlice)) // Sum of floats: 6.6

	// stringのスライスを渡そうとすると、コンパイルエラーになる！
	// stringSlice := []string{"a", "b"}
	// SumNumbers(stringSlice) // Error: string does not implement Number
}

/*
このように、型制約を使うことで、安全に特定の操作（この場合は +）を行えることがコンパイル時に保証されます。
*/
