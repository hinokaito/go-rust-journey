package main

import "fmt"

// このように、ファイルの上に書くと、理解しやすくなります。
// const()で囲む場合、constを省略できます。

const (
	// 円周率(プログラム中で変わることが無い)
	Pi = 3.14
	// アプリケーションのバージョン
	Version = "1.0.0"
)

func main() {
	fmt.Println(Pi)      // 3.14
	fmt.Println(Version) // 1.0.0
}
