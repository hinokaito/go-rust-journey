package main

import (
	"fmt"
	"io"
	"os"
)

// 21章_ioを先に学習してからこのファイルを読む方がいいです。

/*os.Fileとインターフェース
os.Fileは、開かれたファイルディスクリプタ*1と表す構造体です。この*os.Fileは、
io.Readerやio.Write、io.Closer、io.Seekerといった多くの標準インターフェースを実装しています。

これにより、os.Fileはfmt.Fprint(io.Writerを引数に取る)やio.Copy(io.Writer, io.Readerを引数に取る)など、
GoのI/O関連の多くの関数とシームレスに連帯できます


/*1
ファイルディスクリプタ（file descriptor, FD）とは、オペレーティングシステム（OS）がオープン中のファイルやソケット、
パイプなどの入出力リソースを識別するために割り当てる一意の整数値
*/

func main() {
	// os.Stdoutも*os.Fileであり、io.Writerインターフェースを満たす
	// 標準出力に書き込む
	from, err := os.Open("source.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer from.Close()

	to, err := os.Create("destination.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer to.Close()

	// io.CopyはReaderからWriterへデータを効率的にコピーする
	_, err = io.Copy(to, from)
	if err != nil {
		fmt.Println(err)
	}
}
