package main

import (
	"fmt"
	"time"
)

/*タイムゾーンの扱い
time.Now()で取得した時刻には、実行環境のローカルタイムゾーン情報が含まれています。
しかし、サーバーアプリケーションなどでは、タイムゾーンを明示的に扱う必要があります。

・time.LoadLocation(name string): タイムゾーンのデータベースから、指定した名前(例: "Asia/Tokyo", "UTC")の*time.Locationを取得します。

・time.In(loc *Location): time.Timeを指定したタイムゾーンの時刻に変換します。
*/

func main() {
	now := time.Now()
	fmt.Println("ローカル時刻:", now)

	// UTC(協定世界時)の時刻に変換
	utcTime := now.UTC()
	fmt.Println("UTC時刻:", utcTime)

	// 東京のタイムゾーン情報を読み込む
	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("タイムゾーンの読み込みエラー:", err)
		return
	}

	// 現在時刻を東京の時刻として解釈
	tokyoTime := now.In(tokyo)
	fmt.Println("東京時刻:", tokyoTime)

	// 特定の時刻を特定のタイムゾーンで生成
	specificTimeInTokyo := time.Date(2025, 1, 1, 0, 0, 0, 0, tokyo)
	fmt.Println("東京の特定の時刻:", specificTimeInTokyo)
}

/*注意
サーバー環境によってはタイムゾーンデータ(IANA Time Zone Database)がインストールされていない場合、LoadLocationは失敗します。
*/
