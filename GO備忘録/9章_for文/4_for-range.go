package main

import "fmt"

/*
	スライス、配列、マップなどのコレクションの各要素を順番に取り出したい場合に、非常に便利なforループです。
	この形式はGoで最もよく使われるループの一つです。

	基本構文 (スライス/配列):
    for インデックス, 要素 := range コレクション {
		繰り返したい処理
	}

	基本構文 (マップ):
    for キー, 値 := range マップ {
		繰り返したい処理
	}
*/

func main() {
	fruits := []string{"リンゴ", "バナナ", "イチゴ"}

	// スライスの要素をループで取り出す
	for index, value := range fruits {
		fmt.Printf("インデックス: %d, 値: %s\n", index, value)
	}

	// インデックスが不要な場合
	fmt.Println("--- 値だけ ---")
	for _, value := range fruits {
		fmt.Println(value)
	}

	// マップの要素をループで取り出す
	scores := map[string]int{"国語": 85, "数学": 92}
	for subject, score := range scores {
		fmt.Printf("%sは%d点\n", subject, score)
	}
}
