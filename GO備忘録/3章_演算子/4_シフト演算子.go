package main

import "fmt"

const (
	a = 10
	b = 3
)

func main() {
	fmt.Printf("%08b\n", a<<b) // 101000 3ビット左シフト
	fmt.Printf("%08b\n", a>>b) // 0001 3ビット右シフト
}
