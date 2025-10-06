package main

import (
	"fmt"
	"os"
)

/*
プログラムの実行時の振る舞いを外部から制御するための機能があります。

ここではコマンドライン引数を使ってみましょう。

・os.Args: プログラム実行時に渡されたコマンドライン引数を格納する文字列スライスです
・os.Args[0]: プログラム名自身です
*/

// go run 4_コマンドライン引数.go arg1 arg2 で実行してみてください

func main() {
	fmt.Println("プログラム名:", os.Args[0])
	fmt.Println("引数:", os.Args[1:])
}
