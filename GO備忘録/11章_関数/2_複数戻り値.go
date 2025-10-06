package main

import (
	"errors" // エラーを作成するためのパッケージ
	"fmt"
)

// int型の計算結果と、error型の情報を2つ返す
// 型が2つある場合は、()でまとめて型を指定します
func divide(a, b int) (int, error) {
	// もし割る数が0なら
	if b == 0 {
		// ゼロ値と、エラー情報を返す
		return 0, errors.New("0で割ることはできません") // errors.New()は、エラー情報を作成する関数
	}

	// 正常に計算で来た場合
	result := a / b
	// 計算結果と、エラーなし(nil)を返す
	return result, nil
}

func main() {
	// 正常なケース
	result, err := divide(10, 2)
	if err != nil {
		// エラーが発生した場合
		fmt.Println("エラー:", err)
	} else {
		// エラーが発生しなかった場合
		fmt.Println("計算結果:", result)
	}

	// エラーが発生するケース
	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("エラー:", err2)
	} else {
		fmt.Println("計算結果:", result2)
	}
}
