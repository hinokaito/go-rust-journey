package main

import "fmt"

/*
Go言語にはwhile文がありません。ループは全てfor文で行います。
以下のようにすることで、while文のように扱うことができます。
*/

func main() {
	n := 1
	// nが5以下の間、ループを続ける
	for n <= 5 {
		fmt.Println(n)
		n++ // nを1増やす（これを忘れると無限ループになる！）
	}
}
