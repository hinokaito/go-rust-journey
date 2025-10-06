/*
encoding/jsonは、Goのデータ構造と、ウェブAPIなどで広く使われるデータ交換フォーマットであるJSONと相互に変換する標準パッケージです。
Goで、APIサーバーやクライアントを開発する上で、避けては通れない非常に重要なパッケージです。
*/

/*主な機能
encoding/jsonの主な機能は2つ、「マーシャリング」と「アンマーシャリング」です。
1. マーシャリング: Goのデータ構造(主に構造体やマップ)をJSON形式のバイト列([]byte)に変換すること
2. アンマーシャリング: JSON形式のバイト列をGoのデータ構造に変換すること。
*/

/*マーシャリング: Go構造体 → JSON
json.Marshal(v any)関数を使います。Goの構造体を引数に渡すと、JSONに変換された[]byteとerrorが返ってきます
重要なルール:
JSONに変換したい構造体のフィールドは、他のパッケージからアクセスできるように、大文字で始まる必要があります。(publicなフィールド)
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	ID       int
	Name     string
	IsActive bool
}

func main() {
	user := User{
		ID:       1,
		Name:     "Taro Yamada",
		IsActive: true,
	}

	// 構造体をJSONバイト列にマーシャリング
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("JSONマーシャリングに失敗しました: %v", err)
	}

	// バイト列を文字列に変換して出力
	fmt.Println(string(jsonBytes))
}
