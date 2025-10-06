package main

import (
	"fmt"
	"time"
)

/*
Goの時刻フォーマットは、他の言語と異なり、特定の基準時刻(Mon Jan 2 15:04:05 -0700 MST 2006)を使ってレイアウトを指定するのが大きな特徴です。
この覚え方は、「1月2日3字4分5秒2006年」と数字を順に並べたものです。

・time.Format(layout string): time.Timeオブジェクトを、指定したレイアウトの文字列に変換します。
・time.Parse(layout, value string): 指定したレイアウトの文字列をtime.Timeオブジェクトに変換(パース)します。
*/

func main() {
	now := time.Now()

	// "2006-01-02 15:04:05" という形式でフォーマット
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("フォーマット後:", formattedTime)

	// 文字列からtime.Timeへパース
	// レイアウト文字列とパースしたい文字列の形式が完全に一致している必要がある
	timeStr := "2025/10/05 12:30:00"
	parsedTime, err := time.Parse("2006/01/02 15:04:00", timeStr)
	if err != nil {
		fmt.Println("パースエラー:", err)
		return
	}
	fmt.Println("パース後:", parsedTime)
}
