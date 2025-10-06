package main

import (
	"fmt"
	"io" // 次の章で取り扱います
	"os"
)

/*ファイルの読み込み
・os.Open(name string): 読み込み専用でファイルを開きます
・os.ReadFile(name string): ファイル全体を一度に読み込み、その内容をバイトスライスで返します。小さなファイルの読み込みに便利です。
・(*of.File).Read(b []byte): *os.Fileオブジェクトからバイトスライスにデータを読み込みます。
*/

func main() {
	// ファイル全体を一度に読み込む
	data, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("ReadFileエラー:", err)
		return
	}
	fmt.Println("ReadFileの内容:", string(data))

	// ファイルを少しずつ読み込む
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println("Openエラー:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 16) // 16バイトのバッファ
	for {
		n, err := file.Read(buffer)
		if err == io.EOF { // ファイルの終端に達したらループを抜ける
			break
			if err != nil {
				fmt.Println("Readエラー:", err)
				return
			}
			fmt.Print(string(buffer[:n]))
		}
	}
}
