package main

import "fmt"

/*
ゴルーチンがチャネルに値を送りつける場合、受信側は「いつになったらデータが終わるのか」を知る必要がありません。
そのためにcloseとrangeを使います

・close(ch): 送信側が「もうこれ以上このチャネルには値を送りません」と宣言するために使います。
・for range ch: チャネルがcloseされるまで、チャネルから値を延々と受信し続けるためのループ構文です。


重要なルール
・closeは送信側だけが行うべきです
・閉じたチャネルに値を送信しようとするパニック(プログラムのクラッシュ)になります。
・閉じたチャネルから受信しようとすると、バッファに残っている値を受け取った後、その型のゼロ値が即座に返され続けます。
*/

// コード例

// producerはチャネルに値を送信し、最後にチャネルを閉じる
func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("送信: %d\n", i)
		ch <- i
	}
	close(ch) // これが重要
	fmt.Println("チャネルを閉じました")
}

func main() {
	ch := make(chan int)
	go producer(ch)

	// for...rangeを使うと、チャネルがcloseされるまで値を受信し続ける
	// closeされると、このループは自動的に終了する
	for value := range ch {
		fmt.Printf("受信: %d\n", value)
	}

	fmt.Println("処理完了")
}
