package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func main() {
	var operator string
	var a, b int

	for {

		fmt.Printf("数字を空白区切りで入力してください: ")
		fmt.Scan(&a, &b)

		fmt.Printf("四則演算の演算子を入力してください(終了するにはqを入力): ")
		fmt.Scan(&operator)

		switch operator {
		case "+":
			fmt.Printf("加算結果: %d\n", add(a, b))
		case "-":
			fmt.Printf("減算結果: %d\n", sub(a, b))
		case "*":
			fmt.Printf("乗算結果: %d\n", mul(a, b))
		case "/":
			fmt.Printf("除算結果: %d\n", div(a, b))
		case "q":
			return
		default:
			fmt.Println("無効な演算子です")
		}
	}
}
