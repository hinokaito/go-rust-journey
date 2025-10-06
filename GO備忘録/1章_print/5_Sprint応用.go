package main

import "fmt"

func main() {
	/*
		Sprint系は、コンソールに出力する代わりに、整形した結果を文字列として返します。
		ログメッセージを組み立てたり、他の関数に渡す文字列を生成するときに便利です。
	*/
	message := fmt.Sprintf("ユーザーID: %d がログインしました。", 123)
	fmt.Println(message) // "ユーザーID: 123 がログインしました。"

	/*
		必須の方でも学んだように、Sprint系には、
		Sprintln, Sprintf, Sprintの3種類があり、同じ性質です。
	*/
}
