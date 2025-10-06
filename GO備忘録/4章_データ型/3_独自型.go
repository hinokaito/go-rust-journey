package main

import "fmt"

/*
typeキーワードを使って、既存の型を基に自分だけの新しい型を定義できます。
これにより、コードの可読性や安全性を向上させることができます。
*/

// `int`を基に `UserID` という新しい型を定義
type UserID int

// `string`を基に `Email` という新しい型を定義
type Email string

func main() {
	var uid UserID = 101
	var email Email = "test@example.com"

	fmt.Printf("uid: 型=%T, 値=%v\n", uid, uid)
	fmt.Printf("email: 型=%T, 値=%v\n", email, email)

	// var wrongId int = uid // これはエラーになる！ UserID型とint型は別物
}
