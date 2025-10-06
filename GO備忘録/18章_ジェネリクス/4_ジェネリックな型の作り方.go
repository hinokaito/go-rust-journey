package main

import "fmt"

/*
関数だけでなく、struct などの型自体もジェネリックにできます。
これにより、特定の型に依存しないデータ構造（スタック、キュー、連結リストなど）を簡単に作れます
*/

// T型のデータを保持するStackを定義
type Stack[T any] struct {
	items []T
}

// Push メソッド (レシーバもジェネリックになる)
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop メソッド
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T // T型のゼロ値を返す
		return zero, false
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

func main() {
	// intを扱うスタックを作成
	intStack := &Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	v, _ := intStack.Pop()
	fmt.Println(v) // 20

	// stringを扱うスタックを作成
	stringStack := &Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")
	s, _ := stringStack.Pop()
	fmt.Println(s) // world
}

/*
Stack[int] と Stack[string] は、同じ Stack の定義から作られていますが、それぞれが型安全な、独立したデータ構造として機能します。
*/
