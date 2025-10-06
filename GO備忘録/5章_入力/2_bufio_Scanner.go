package main

import (
	"bufio" // バッファリングされた入力を読み込むためのパッケージ
	"fmt"
	"os"      // ファイルの読み込みや書き込みを行うためのパッケージ
	"strconv" // 文字列を数値に変換するためのパッケージ
)

/*
おっと3つも知らないパッケージが出てきましたね。
それぞれ難しいものではないので軽く知っておきましょう。
共通して言えることですが、概要のみ知り、必要になったら学ぶのがタイパ良いです。
*/

// 標準入力を読み取るためのスキャナーをグローバルで用意
// 関数外のとき、:=は使えません。
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	// --- 文字列を一行読み込む ---
	fmt.Println("フルネームで入力してください: ")
	// scanner.Scan()で一行読み込む
	scanner.Scan()

	// scanner.Text()で読み込んだ文字列を取得
	fullName := scanner.Text()

	// --- 数値を一行読み込む ---
	fmt.Println("年齢を入力してください: ")
	scanner.Scan()

	// 読み込んだ文字列を Atoi (As To Int)で数値に変換
	// エラーハンドリングのために2つ値が返されますが、今は使わないので、_(仮の値)で受け取ります。
	age, _ := strconv.Atoi(scanner.Text())

	// --- 読み込んだデータを表示 ---
	fmt.Printf("こんにちは、%sさん(%d歳)\n", fullName, age)
}
