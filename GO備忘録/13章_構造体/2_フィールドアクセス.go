package main

import "fmt"

/*
構造体のフィールドにアクセスしたり、値を変更したりするには、ドット(.)を使います。
*/

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{Name: "Charlie", Age: 40}

	// フィールドの値を取得
	fmt.Println("名前:", user.Name) // 名前: Charlie

	// フィールドの値を変更
	user.Age = 41
	fmt.Println("変更後の年齢:", user.Age) // 変更後の年齢: 41
	fmt.Println("変更後のユーザー:", user)   // 変更後のユーザー: {Charlie 41}
}
