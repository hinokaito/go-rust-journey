package main

import "fmt"

/*
deferは、Goの強力な機能の1つで、関数の終了時に実行したい処理を予約するためのキーワードです。
後片付け処理を簡潔かつ確実に記述できるため、Goのコードでは頻繁に利用されます。

基本的な仕組みとしては、deferキーワードの後に続く関数呼び出しを、それが記述された関数が終了する直前まで遅延させます。
*/

func main() {
	defer fmt.Println("world") // この行の実行が遅延される

	fmt.Println("hello")
}

/*実行結果
hello
world

main関数の中でdefer fmt.Println("world")が先に書かれていますが、実行されるのはmain関数が終わる直前の時点です。そのため、helloが先に表示されます。
*/
