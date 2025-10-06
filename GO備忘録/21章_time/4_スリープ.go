package main

import (
	"fmt"
	"time"
)

/*スリープ
time.Sleep(d Duration): 指定したDurationだけ、現在のゴルーチンを停止させます。
*/

func main() {
	fmt.Println("処理を開始します")
	time.Sleep(2 * time.Second) // 2秒間停止
	fmt.Println("2秒経過しました")
}
