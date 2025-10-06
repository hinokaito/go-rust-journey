package main

import "fmt"

/*
switch文は、条件分岐を行うための構文です。
if文と似ていますが、以下のポイントで使い分けるとよいです。

・条件が複雑な時　→　if文
・条件が1つか2つのとき　→　if文
・評価対象がバラバラの値や式のとき　→　if文

・分岐が多い時　→　switch文
・型ごとの処理分け　→　switch文
・条件式の見通しを良くしたいとき　→　switch文
・同じ値を複数の候補と比較するとき　→　switch文
*/

func main() {
	var signal string
	signal = "red"

	switch signal {
	case "red":
		fmt.Println("止まれ")
	case "yellow":
		fmt.Println("注意")
	case "blue":
		fmt.Println("進め")
	default:
		fmt.Println("信号機が故障しています")
	}

	// if文で書いてみましょう
	if signal == "red" {
		fmt.Println("止まれ")
	} else if signal == "yellow" {
		fmt.Println("注意")
	} else if signal == "blue" {
		fmt.Println("進め")
	} else {
		fmt.Println("信号機が故障しています")
	}

	// 明らかにswitch文の方が見やすいですね
}
