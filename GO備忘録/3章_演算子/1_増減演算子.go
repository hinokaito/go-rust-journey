package main

import "fmt"

func main() {
	a := 10
	a++            // a = a + 1と同義
	fmt.Println(a) // 11
	a--            // a = a - 1と同義
	fmt.Println(a) // 10
}
