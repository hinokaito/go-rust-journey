package main

import (
	"fmt"
	"os"
)

/*ディレクトリ操作
・os.Mkdir(name string, perm fs.FileMode): 新しいディレクトリを作成します。
・os.MkdirAll(path string, perm fs.FileMode): ネストしたディレクトリを作成します。
・os.ReadDir(name string): ディレクトリの内容を読み込み、ディレクトリエントリのスライスを返します。
・os.Remove(name string): ファイルまたは空のディレクトリを削除します。
・os.RemoveAll(path string): ディレクトリとその中身を全て再帰的に削除します。非常に強力で危険です。
*/

// ディレクトリ操作をしてみましょう

func main() {
	// 単一のディレクトリを作成
	err := os.Mkdir("testdir", 0755)
	if err != nil {
		fmt.Println("Mkdirエラー:", err)
	} else {
		fmt.Println("testdirディレクトリを作成しました")
	}

	// ネストしたディレクトリを作成
	err = os.MkdirAll("parent/child/grandchild", 0755)
	if err != nil {
		fmt.Println("MkdirAllエラー:", err)
	} else {
		fmt.Println("parent/child/grandchildディレクトリを作成しました")
	}

	// ディレクトリの内容を読み込み
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("ReadDirエラー:", err)
		return
	}

	fmt.Println("現在のディレクトリの内容:")
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("📁 %s (ディレクトリ)\n", entry.Name())
		} else {
			fmt.Printf("📄 %s (ファイル)\n", entry.Name())
		}
	}

	// 空のディレクトリを削除
	err = os.Remove("testdir")
	if err != nil {
		fmt.Println("Removeエラー:", err)
	} else {
		fmt.Println("testdirディレクトリを削除しました")
	}

	// ネストしたディレクトリを再帰的に削除
	err = os.RemoveAll("parent")
	if err != nil {
		fmt.Println("RemoveAllエラー:", err)
	} else {
		fmt.Println("parentディレクトリとその中身を削除しました")
	}
}
