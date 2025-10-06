package main

import "fmt"

/*
Goは自動でbreakしますが、
あるcaseの処理を実行した後、次のcaseの処理も続けて実行したい場合があります。
その場合はfallthroughを使います。

注意:
fallthroughは次のcaseの条件を評価せずに、無条件でその処理ブロックにジャンプします。
使う場面は限定的で注意深く利用する必要があります
*/

func main() {
	num := 2

	switch num {
	case 1:
		fmt.Println("1です")
	case 2:
		fmt.Println("2です")
		fallthrough
	case 3:
		fmt.Println("3です")
	case 4:
		fmt.Println("4です")
	}
}
