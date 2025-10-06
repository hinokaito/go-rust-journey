/*
実際にコードを見てみましょう。
「鳴く(Cry)」という振る舞いをインターフェースとして定義し、DogとCatという2つの構造体に実装させてみます。
*/
package main

import "fmt"

// Animalインターフェースを定義
// Cry()というメソッドを持つことを契約する
type Animal interface {
	Cry() string
}

// Dog構造体
type Dog struct {
	Name string
}

// Dog型にCry()メソッドを実装(Animalインターフェースの契約を満たす)
func (d Dog) Cry() string {
	return "ワン！"
}

// Cat構造体
type Cat struct {
	Name string
}

// Cat型にCry()メソッドを実装(こちらも契約を満たす)
func (c Cat) Cry() string {
	return "ニャー！"
}

// Animalインターフェースを引数に取る関数
// この関数は、具体的な型(Dog, Cat)を知らない。
// Cry()メソッドを持っていることだけを知っている。
func makeSound(a Animal) {
	fmt.Println(a.Cry())
}

func main() {
	// DogとCatのインスタンスを作成
	pochi := Dog{Name: "ポチ"}
	tama := Cat{Name: "タマ"}

	// DogもCatも、Animalインターフェースを満たしているので、
	// makeSound関数に渡すことができる
	makeSound(pochi) // ワン！
	makeSound(tama)  // ニャー！

	// インターフェース型のスライスを作ることもできる
	animals := []Animal{pochi, tama}
	for _, animal := range animals {
		fmt.Println(animal.Cry())
	} // ワン！
	// ニャー！
}

/*
makeSound関数は、具体的な型(Dog, Cat)ついては一切感知していません。
ただ「Animalという契約を満たすものなら何でも受け入れる」という姿勢です。
これにより、将来BirdやCowといった新しい構造体を追加しても、makeSound関数を一切変更する必要はありません。
*/
