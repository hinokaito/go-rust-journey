package main

import "fmt"

// for文については、8章で詳しく解説します。
// そこではこのコードを再利用するので飛ばしても構いません。

func main() {
	users := map[int]string{
		101: "山田",
		102: "佐藤",
		103: "鈴木",
	}

	// for key, value := range myMap
	for id, name := range users {
		fmt.Printf("ID: %d, 名前: %s\n", id, name)
	}
	// 注意：実行するたびに表示される順序が変わる可能性があります！

	/*
		結果
		ID: 101, 名前: 山田
		ID: 102, 名前: 佐藤
		ID: 103, 名前: 鈴木
	*/
}
