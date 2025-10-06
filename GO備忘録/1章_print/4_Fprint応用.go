package main

import (
	"fmt"
	"os" // 標準入出力のためのパッケージ このような記述で複数インポートできる
)

/*
	fmtパッケージには、
	Print系以外に、Fprint系と、Sprint系の関数があります。
	ここではFprintを説明します
*/

func main() {
	// 標準エラー出力に書き出す
	fmt.Fprintln(os.Stderr, "これはエラーメッセージです。") // これはエラーメッセージです。
	/*なぜこれが必要？
	出力先を標準出力以外に指定できる！
	例えば、ファイルに書き出すときにも使う

	必須の方でも学んだように、Fprint系には、
	Fprintln, Fprintf, Fprintの3種類があり、同じ性質です。
	*/
}
