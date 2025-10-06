package main

import "fmt"

/*
引数の数を固定せず、任意の数の引数を受け取れる関数を定義することもできます。
fmt.Printfを思い出してください。好きなだけ引数を渡せますよね。それと同じように、可変長引数を使うと、任意の数の引数を受け取ることができます。
実装方法は比較的簡単で、引数の型の前に...を付けるだけです。
*/

func sumAll(numbers ...int) int {
	// numbersはスライス([]int)として扱われる
	fmt.Printf("受け取った引数: %v\n", numbers)

	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// 1つの引数を渡す
	fmt.Println("合計:", sumAll(1))

	// 2つの引数を渡す
	fmt.Println("合計:", sumAll(1, 1))

	// 5つの引数を渡す
	fmt.Println("合計:", sumAll(1, 1, 1, 1, 1))

	// スライスを渡す場合は、最後に ... を付ける
	nums := []int{1, 10, 100}
	fmt.Println("合計:", sumAll(nums...))
}
