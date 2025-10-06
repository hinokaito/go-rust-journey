package main

import "fmt"

func main() {
	a, b := 1, 2

	fmt.Println(a) // 1
	fmt.Println(b) // 2

	a, b = b, a

	fmt.Println(a) // 2
	fmt.Println(b) // 1
}
