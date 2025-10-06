package main

// 実際に、エラーを作成し、それをハンドリングするコードを見てみましょう。

import (
	"errors" // シンプルなエラーを作成するパッケージ
	"fmt"
	"log" // ログを出力し、プログラムを終了するパッケージ
)

// 年齢を検証し、有効なら(int, nil)を、無効なら(0, error)を返す関数
func validateAge(age int) (int, error) {
	if age < 0 {
		// fmt.Errorf を使うと、Printfのようにフォーマットされたエラー文字列を作れる
		return 0, fmt.Errorf("年齢が不正です: %d", age)
	}

	if age < 18 {
		// error.New は、固定の文字列からシンプルなエラーを作る
		return 0, errors.New("未成年です")
	}

	// エラーがない場合は、nilを返す
	return age, nil
}

func main() {
	// 1. 成功するケース
	age1, err1 := validateAge(30)
	if err1 != nil {
		fmt.Printf("エラーが発生しました: %v\n", err1)
	} else {
		fmt.Printf("認証成功: %d歳\n", age1)
	}

	// 2. 失敗するケース(未成年)
	age2, err2 := validateAge(15)
	if err2 != nil {
		fmt.Printf("エラーが発生しました: %v\n", err2)
	} else {
		fmt.Printf("認証成功: %d歳\n", age2)
	}

	// 3. 失敗するケース(不正な値)
	age3, err3 := validateAge(-5)
	if err3 != nil {
		// エラーが発生したら、ログを出力してプログラムを終了させる(log.Fatal)
		log.Fatalf("致命的なエラーが発生しました: %v", err3)
	} else {
		fmt.Printf("認証成功: %d歳\n", age3)
	}

	/*
		認証成功: 30歳
		エラーが発生しました: 未成年です
		2025/10/04 17:42:00 致命的なエラーが発生しました: 年齢が不正です: -5
		exit status 1
	*/
}

/*
一見冗長と感じるかもしれませんが、プログラムのどこでエラーが起こりうるかを常に意識させ、
結果として非常に堅牢で信頼性の高いソフトウェアにつながります。
*/
