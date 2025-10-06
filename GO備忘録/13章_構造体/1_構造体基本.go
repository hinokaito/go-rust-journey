package main

import "fmt"

/*
構造体とは、複数の異なる型のデータ(フィールドと呼ぶ)を1つに束ねて、新しいデータ型を定義するための仕組みです
例えば「ユーザー」というデータを扱いとき、
名前(string), 年齢(int), 有効フラグ(bool) といった複数の情報をまとめて管理できます。


なぜ構造体を使うのか？

・関連データの集約: userName, userAge, userActiveのように変数をばらばらに管理するのではなく、Userという1つのまとまりとして扱え、コードが整理されます。

・可読性の向上: user.Name, user.Age のように、何の情報にアクセスしているのかが一目瞭然になります。

・関数の引数や戻り値として便利: 多くの情報を1つの構造体に入れてやり取りできるため、関数がシンプルになります。

このように構造体を定義できます。
type 構造体名 struct {
	フィールド名1 型1
	フィールド名2 型2
	...
}

インスタンス化には様々な方法がありますが、最も推奨される方法はこのように書きます
インスタンス名 := 構造体名{
	フィールド名1: 値1,
	フィールド名2: 値2,
	...
}*/

// 実際にコードを書いてみましょう

// 構造体の定義
type User struct {
	Name     string
	Age      int
	IsActive bool
}

func main() {
	// 1. フィールド名を指定してインスタンス化(最も推奨される方法)
	// 順番を気にする必要がなく、どのフィールドに何を設定しているか明確
	user1 := User{
		Name:     "Alice",
		Age:      30,
		IsActive: true,
	}
	fmt.Println("User1:", user1) // User1: {Alice 30 true}

	// 2. フィールド名を省略してインスタンス化(非推奨)
	// 定義の順番通りに値を書く必要があるため、変更に弱い
	user2 := User{"Bob", 25, false}
	fmt.Println("User2:", user2) // User2: {Bob 25 false}

	// 3. ゼロ値の構造体を作成
	// 各フィールドがゼロ値(stringは"", intは0, boolはfalse)で初期化される
	var user3 User
	fmt.Println("User3(ゼロ値):", user3) // User3(ゼロ値): { 0 false}
}
