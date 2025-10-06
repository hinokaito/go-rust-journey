// 言葉だけでは難しいので実際のコードで動きを見ていきましょう
package main

import "fmt"

func main() {
	// 1. 通常の変数を宣言
	num := 42
	fmt.Printf("numの値: %d\n", num)

	// 2. '&'をつかって、num変数のアドレスをポインタ変数ptrに格納
	// ptrの型は「intへのポインタ」を意味する *int 型になる
	var ptr *int = &num

	fmt.Printf("numのアドレス(ptrの値): %p\n", ptr)

	// 3. '*'をつかって、ptrが指すアドレスにある値を取得
	fmt.Printf("ptrが指す先の値: %d\n", *ptr)

	// 4. ポインタを通じて、元の変数numの値を変更する
	fmt.Printf("\n--- ポインタ経由で値を変更 ---")
	*ptr = 100

	fmt.Printf("ポインタ経由で変更した後のptrが指す先の値: %d\n", *ptr)
	fmt.Printf("元の変数numの値も変わっている: %d\n", num) // numの値が100になっている
}
