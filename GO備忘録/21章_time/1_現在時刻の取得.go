package main

import (
	"fmt"
	"time"
)

/*
timeパッケージは、Goで時刻の取得や計算、時間の計測、タイマー処理など、時間に関するあらゆる操作を行うための標準パッケージです
正確性と使いやすさを両立しており、Goの重要な機能の1つです。
*/

/*時間の表現
timeパッケージには、中心となる2つの型があります。

1. time.Time: ある特定の「瞬間」を表すための構造体です。年月日時分秒、ナノ秒、そしてタイムゾーンの情報を持っています。

2. time.Duration: 2の瞬間の「間隔」を表すための型です。ナノ秒単位の整数として表現されます。
   time.Hourやtime.Minuteといった便利な定数が用意されています。const
*/

// 現在時刻の取得
// time.Now()で現在のローカル時刻をtime.Time型で取得します。

func main() {
	// 現在時刻を取得
	now := time.Now()
	fmt.Println("現在時刻:", now)

	// 年月日などを個別に取得
	fmt.Printf("年: %d, 月: %d, 日: %d", now.Year(), now.Month(), now.Day())
	fmt.Printf("時: %d, 分: %d, 秒: %d", now.Hour(), now.Minute(), now.Second())
	fmt.Printf("ナノ秒: %d", now.Nanosecond())
	fmt.Printf("タイムゾーン: %s", now.Location())

	// 特定の日時を表すtime.Timeを作成
	specificTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	fmt.Println("特定の日時:", specificTime)
}
