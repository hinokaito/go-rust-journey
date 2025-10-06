package main

/*
1.select文とは
select文は、複数のチャネル操作を同時に持つための仕組みで、「チャネル版のswitch文」と考えると分かりやすいです

2.特徴
・複数のcaseに書かれたチャネル送受信のうち、最初に準備ができたものを実行します。
・もし複数のcaseが同時に準備できた場合は、ランダムにいずれかの1つが選ばれます。
・default節を追加すると、どのcaseも準備できていないばあいに即座に実行され、select文がブロックしなくなります。

3.どういうときに使うの？
・複数のゴルーチンからの結果を待ち受けるとき。
・タイムアウト処理を実装するとき。
・ゴルーチンに中断命令を送るとき。
*/

// コード例

import (
	"fmt"
	"time" // 遅延するために使います
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// 1秒後にc1にメッセージを送信するゴルーチン
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "c1からのメッセージ"
	}()

	// 2秒後にc2にメッセージを送信するゴルーチン
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "c2からのメッセージ"
	}()

	// 2回ループして、両方のメッセージを受け取る
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("受信:", msg1)
		case msg2 := <-c2:
			fmt.Println("受信:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("タイムアウトしました")
		}
	}
}

/*実行結果
受信: c1からのメッセージ
受信: c2からのメッセージ
*/
