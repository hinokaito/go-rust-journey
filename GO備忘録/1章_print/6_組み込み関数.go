package main

/*
	組み込み関数printとprintlnについて説明します。
	これらの関数は、主にデバッグやGo自信の開発に用意されたもので、
	本番環境のコードで使うことは公式に推奨されていません。

	以下の特徴を持ちます
	・パッケージ不要: importしなくても使える
	・低レベル: Goのランタイムに直接アクセスするため、fmtより高速
	・機能制限：
	1.書式設定(%sなど)は利用できない。 2.出力先は常に標準エラー出力(stderr)に固定
	3.fmtほど賢くなく、ポインタの参照先を表示してくれないと挙動が異なる場合がある

*/

func main() {
	lang := "Go"
	num := 100

	// 組み込みのprintとprintlnを使ってみる
	print("This is a built-in print. lang:", lang, " num:", num, "\n") // This is a built-in print. lang:Go num:100
	println("This is a built-in println.", "lang:", lang, "num:", num) // This is a built-in println. lang:Go num:100
}

// 基本、fmtを使えばいいです..
