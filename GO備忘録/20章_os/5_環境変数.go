package main

import (
	"fmt"
	"os"
)

/*環境変数
・os.Getenv(key string): 指定されたキーの環境変数の値を取得します。存在しない場合は空文字列を返します。
・os.Setenv(key, value, string): 環境変数を設定します
・os.LookupEnv(key string): 環境変数の値と、その変数が存在するかどうかをbool値で返します。
　変数が未設定なことと、空文字列が設定されていることを区別したい場合に便利です。
*/

func main() {
	// HOME環境変数を取得
	home := os.Getenv("HOME")
	fmt.Println("HOME環境変数:", home)

	// 環境変数を設定
	os.Setenv("MY_VAR", "my_value")

	// 設定した環境変数を取得
	val, ok := os.LookupEnv("MY_VAR")
	if ok {
		fmt.Println("MY_VAR:", val)
	}
}
