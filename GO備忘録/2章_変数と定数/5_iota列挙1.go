package main

import "fmt"

/*
	iota(イオタ)とは、const宣言ブロック内でのみ使える特殊な定数です。
	0から始まり、1行ごとに1ずつインクリメント(増加)される便利な仕組みです

	主に、連番の定数を定義するのに使われます
*/

const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

func main() {
	fmt.Printf("%d", Sunday)
	fmt.Printf("%d", Monday)
	fmt.Printf("%d", Tuesday)
	fmt.Printf("%d", Wednesday)
	fmt.Printf("%d", Thursday)
	fmt.Printf("%d", Friday)
	fmt.Printf("%d", Saturday)
	// 0123456
}
