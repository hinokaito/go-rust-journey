package main

import "fmt"

/*
他の言語にある condition ? true_val : false_val のような三項演算子は、Goには存在しません
これはGoの設計哲学である、
「読みやすさは重要であるが、書きやすさよりも重要である」
「物事をおこなう方法は1つであるべき」
という考えに基づいています。
if-elseを使えば誰が読んでも処理の流れが明確になる、というのがGoのスタイルです。
*/

func main() {
	// 他の言語: const result = score > 60 ? "合格" : "不合格";

	// Goの場合: if-elseではっきりと書く
	score := 80
	var result string
	if score > 60 {
		result = "合格"
	} else {
		result = "不合格"
	}
	fmt.Println(result) // 合格
}
