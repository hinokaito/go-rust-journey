/*
アンマーシャリング: JSON → Go構造体
json.Unmarshal(data []byte, v any)関数を使います。JSONのバイト列と、変換結果を格納するGoのデータ構造のポインタを引数に渡します。
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
	jsonStr := `{"ID:2,"Name":"Jiro Yamada","IsActive":true}`
	jsonBytes := []byte(jsonStr)

	var user User

	// JSONバイト列を構造体にアンマーシャリング
	// 第2引数にはポインタを渡す必要がある
	err := json.Unmarshal(jsonBytes, &user)
	if err != nil {
		log.Fatalf("JSONアンマーシャリングエラー: %v", err)
	}

	fmt.Printf("ID: %d, Name: %s, Active: %v\n", user.ID, user.Name, user.IsActive)
}
