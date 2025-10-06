package main

import "fmt"

/*LIFO(後入れ先出し)の実行順序 stack
1つの関数内複数のdefer文が使われた場合、それらはスタックに積まれます。
つまり、最後にdeferで指定した処理から順に実行されます。
*/

func main() {
	fmt.Println("カウント")

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("完了")
}

/*実行結果
カウント
完了
3
2
1
*/

// この性質は、リソースを順番に確保し、それを逆順で開放するような処理を書く際に非常に便利です。
