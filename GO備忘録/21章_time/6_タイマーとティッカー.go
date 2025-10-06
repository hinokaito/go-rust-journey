package main

import (
	"fmt"
	"time"
)

/*タイマーとティッカー
time.Sleepは単純な待機ですが、Goの並行処理と組み合わせることで、より高度な時間制御が可能になります。

・time.NewTimer(d Duration) / time.After(d Duration):
	・Timerは、指定された時間が経過した後に、一度だけ自身のチャネルCに現在時刻を送信するオブジェクトです。
	・time.Afterはtime.NewTimer(d).Cとほぼ同じで、一度きりのイベントを持つ場合に便利なヘルパー関数です。

・time.NewTicker(d Duration):
	・Tickerは、指定した時間間隔で、繰り返しチャネルＣに時刻を送信し続けるオブジェクトです。
	・使い終わったらStop()メソッドを呼んでリソースを開放する必要があります。

・time.Tick(d Duration):
	・NewTickerを使って作成したTickerオブジェクトのCチャネルにアクセスするのと同じ効果があります。
	・しかし、内部的にはTickerを使っているので、Tickerオブジェクトを使うよりも少し効率が悪いです。
	・これは、時間が経過するたびに新しいTickerオブジェクトを作成するのではなく、既存のオブジェクトを使い回すためです。
*/

func main() {
	// ---Timerの使用例---
	fmt.Println("Timerを開始します。2秒後にメッセージが表示されます。")
	timer := time.NewTimer(2 * time.Second)
	<-timer.C // チャネルから受信するまでブロックされる
	fmt.Println("Timer: 2秒経過しました。")

	// ---Tickerの使用例---
	fmt.Println("\nTickerを開始します。1秒ごとにメッセージが5回表示されます。")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // 関数終了時にtickerを停止

	count := 0
	for t := range ticker.C {
		fmt.Println("Ticker:", t.Format("15:04:05"))
		count++
		if count >= 5 {
			break
		}
	}
	fmt.Println("Ticker終了しました。")
}

// これらの機能は、select文と組み合わせることで、より高度なタイムアウト処理やイベント駆動のプログラミングが可能になります。
