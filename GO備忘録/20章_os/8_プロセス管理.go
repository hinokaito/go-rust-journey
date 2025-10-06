// go run 8_プロセス管理.go

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall" // シグナルを扱うためのパッケージ
	"time"
)

/*Goプログラムから他のプロセスを起動したり、シグナルを扱ったりできます。

・プロセスの起動:
	・os.StartProcess(name string, argv []string, attr *ProcAttr): 新しいプロセスを開始するための低レベルAPI
	  より高レベルで使いやすい、os/execパッケージの利用が一般的です。

・プロセスの情報:
	・os.Getpid(): 現在のプロセスのプロセスID(PID)を取得します。
	・os.Getppid(): 親プロセスのPIDを取得します
	・os.FindProcess(pid int): PIDから*os.Processオブジェクトを取得します。

・プロセスの終了:
	・os.Exit(code int): プログラムを即座に終了させます。defer文は実行されないことに注意です。

・シグナルハンドリング:
	OSからのシグナル(例: Ctrl+C)をプログラムで捕捉して、特別な処理(例: 安全なシャットダウン)を行うことができます。os/signalパッケージを使います。
*/

func main() {
	// シグナルを受け取るためのチャネルを作成
	sigChan := make(chan os.Signal, 1)

	// SIGINT(Ctrl+C)とSIGTERM(プログラム終了)を受け取るように設定
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("プロセスを開始しました。終了するにはCtrl+Cを押してください。")

	// ゴルーチンでシグナルを待つ
	go func() {
		sig := <-sigChan
		fmt.Printf("\nシグナル %v を受信しました。シャットダウンします\n", sig)
		// ここでクリーンアップ処理を行う
		time.Sleep(2 * time.Second) //疑似クリーンアップ処理
		os.Exit(0)
	}()

	// メインの処理(ここでは無限ループで待機)
	for {
		fmt.Println(".")
		time.Sleep(500 * time.Millisecond)
	}
}
