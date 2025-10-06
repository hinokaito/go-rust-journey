package main

import (
	"fmt"
	"os"
)

/*
osパッケージは、Goプログラムがオペレーティングシステム(OS)の機能と対話するための中心的な役割を担います。
ファイル操作、プロセス管理、環境変数の取得など、OSレベルの機能にアクセスするための基本的なインターフェースを提供します。

しばらく、Go言語の標準ライブラリの各パッケージを学習していきます。
*/

/*ファイルの作成と書き込み
・os.Create(name string): 新しいファイルを作成します。既にファイルが存在する場合は、中身を空にします(トランケート)。書き込み可能な*os.Fileを返します。

・os.WriteFile(name string, data []byte, perm fs.FileMode): ファイルに倍とスライスを直接書き込みます。
　ファイルのオープン、書き込み、クローズを一度に行う便利な関数です。

・(*os.File).WriteString(s string): *os.Fileオブジェクトに文字列を書き込みます。}
*/

// ファイル操作と書き込みをしてみましょう

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("ファイル作成エラー:", err)
		return
	}
	defer file.Close() // 関数終了時にファイルを必ず閉じる。

	// 文字列を書き込む
	_, err = file.WriteString("Hello, Go OS Library!\n")
	if err != nil {
		fmt.Println("ファイル書き込みエラー:", err)
	}

	// バイトスライスを直接書き込む
	data := []byte("This is direct write.\n")
	err = os.WriteFile("direct.txt", data, 0644) // 0644はファイルのパーミッション
	if err != nil {
		fmt.Println("WriteFileエラー:", err)
	}
}
