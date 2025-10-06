package main

import "fmt"

func main() {
	// if 変数名 := 初期化; 条件式 {
	if score := 80; score >= 60 {
		fmt.Println("合格です")
	} // 合格です
	// ifブロック内でしか使わない変数を扱いたいときに便利です。
}
