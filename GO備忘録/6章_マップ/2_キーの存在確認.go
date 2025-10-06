package main

import "fmt"

// if文については、続く7章で詳しく解説します。
// そこではこのコードを再利用するので飛ばしても構いません。

func main() {
	scores := map[string]int{
		"国語": 85,
		"数学": 92,
	}

	// 存在しないキー "理科" を取得してみる
	scienceScore := scores["理科"]
	fmt.Printf("理科の点数（存在しない場合）: %d\n", scienceScore) // 理科の点数（存在しない場合）: 0

	// 2つの戻り値を使って、キーの存在をチェックする
	// value, ok := map[key]
	// okには、キーが存在すればtrue、存在しなければfalseが入る
	socialScore, ok := scores["社会"]
	if ok {
		fmt.Printf("社会の点数: %d\n", socialScore)
	} else {
		fmt.Println("社会のデータはありません。")
	}

	// 値は不要で、存在チェックだけしたい場合はブランク識別子 `_` を使う
	_, ok = scores["国語"]
	if ok {
		fmt.Println("国語のデータは存在します。")
	}
}
