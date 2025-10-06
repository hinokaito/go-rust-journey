package main

import (
	"log"
	"os"
)

/*カスタムロガーの作成
パッケージレベルの関数（log.Printなど）は便利ですが、これらはすべて単一のグローバルなロガーを共有しています。
ログの出力先をファイルに変えたり、フォーマットを分けたりしたい場合は、自分でLoggerインスタンスを作成します。

log.New(out io.Writer, prefix string, flag int) を使って、新しいロガーを作成できます。

    ・out: ログの出力先です。os.Stderrやos.Stdout、*os.Fileなど、io.Writerインターフェースを満たすものなら何でも指定できます。

    ・prefix: 各ログ行の先頭に付与される文字列です。[INFO]や[ERROR]といったログレベルを示すのによく使われます。

    ・flag: ログのフォーマットを制御する定数です。|（ビットOR）で複数指定できます


	主なフラグ:
定数	　　　　　　　説明	                                 例
log.Ldate	　　　　日付（年/月/日）	                     2025/10/05
log.Ltime	　　　　時刻（時:分:秒）						 12:09:18
log.Lmicroseconds  マイクロ秒								12:09:18.123456
log.Llongfile	   フルファイルパスと行番号					 /path/to/your/main.go:10
log.Lshortfile	   ファイル名と行番号						 main.go:10
log.LUTC	       UTC（協定世界時）を使用
log.LstdFlags	   log.Ldate | log.Ltime のデフォルト設定	2025/10/05 12:09:18

*/

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	// ログファイルを開く
	// os.O_APPEND: 追記モード, os.O_CREATE: なければ作成, os.O_WRONLY: 書き込み専用
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// INFOレベルのロガーを作成
	InfoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)

	// ERRORレベルのロガーを作成 (出力先はファイルと標準エラー出力の両方)
	ErrorLogger = log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Llongfile)
}

func main() {
	InfoLogger.Println("アプリケーションを開始しました。")

	_, err := os.Open("non-existent-file.txt")
	if err != nil {
		ErrorLogger.Printf("致命的なエラーが発生: %v", err)
	}

	InfoLogger.Println("アプリケーションを終了します。")
}
