package main

import (
	"fmt"
	"os"
)

/*ファイルのメタデータ(サイズ、変更日時、パーミッションなど)を扱います。
・os.Stat(name string): ファイルの情報をos.FileInfoインターフェ―スとして取得します。
  os.FileInfoからはファイル名、サイズ、モード(パーミッションやディレクトリかどうか)、最終更新日時などが取得できます。

・os.Chmod(name string, mode FileMode): ファイルのパーミッションを変更します。
・os.Chown(name string, uid, gid int): ファイルの所有者(ユーザーIDとグループID)を変更します。
*/

func main() {
	info, err := os.Stat("hello.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("ファイルが存在しません")
		} else {
			fmt.Println("エラー:", err)
		}
		return
	}

	fmt.Println("ファイル名:", info.Name())
	fmt.Println("サイズ:", info.Size(), "bytes")
	fmt.Println("パーミッション:", info.Mode())
	fmt.Println("ディレクトリ？:", info.IsDir())
	fmt.Println("最終更新日時:", info.ModTime())

	// パーミッションを変更(例: 実行権限を付与)
	err = os.Chmod("hello.txt", 0755)
	if err != nil {
		fmt.Println("Chmodエラー:", err)
	}
}
