package main

import "fmt"

// 値を埋め込んで整形された文章を作りたいときに使う。
// 引数間スペースなし,自動改行なし,フォーマット指定なし。

func main() {
	// 複数の引数を渡しても、スペースや改行は入らない
	fmt.Print("Hello", "Golang!") // Golangisfun!

	// Printlnと違い、自動で改行されない
	fmt.Print("Golang", "is", "fun!") // HelloGolang!Golangisfun!

	// 改行 (他のPrintメソッドも同様に\nは使える)
	fmt.Print("改行しますよ\n")
	fmt.Print("改行されましたよ")
	/*
		出力結果:
		HelloGolang!Golangisfun!改行しますよ
		改行されましたよ
	*/
}
