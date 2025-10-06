package main

import (
	"log"
	"os"
)

/*
Goのlogパッケージは、シンプルで手軽に使えるロギング機能を提供します。
複雑な設定なしに、標準エラー出力に日時やファイル名を含むログメッセージを出力できます。
*/

/*基本的なロギング関数
・log.Print(v ...any) / log.Println(v ...any) / log.Printf(format string, v ...any):
  fmtパッケージの同名関数と似ていますが、出力の先頭に日時情報が自動的に浮揚されます。これらが最も一般的に使われるログ出力関数です。

・log.Fatal(v ...any) / log.Fatalln(v ...any) / log.Fatalf(format string, v ...any):
  Print系の関数と同様にログを出力した後、os.Exit(1)を呼び出してプログラムを即座に終了させます。
  致命的なエラーが発生し、処理の続行が不可能な場合に使用します。

・log.Panic(v ...any) / log.Panicln(v ...any) / log.Panicf(format string, v ...any):
  ログを出力した後、panic()を呼び出します。これによりプログラムはパニック状態になりますが、
  deferとrecoverによって復帰処理が可能です。Fatalと違い、クリーンアップ処理を行うチャンスがあります。
*/

func main() {
	// 最も基本的なログ出力
	log.Println("これは通常のログメッセージです。")
	log.Printf("ユーザーID: %d がログインしました。", 123)

	// ファイルを開いてみる
	_, err := os.Open("non-existent-file.txt")
	if err != nil {
		// エラーが発生し、処理を続行できない場合
		log.Fatalf("ファイルのオープンに失敗しました: %v", err)
	}

	// この行は実行されない
	log.Println("プログラムの終端")
}

/*デフォルトロガーの出力先
デフォルトでは、logパッケージは標準エラー出力(os.Stderr)にログを出力します。
これは、プログラムの正常な出力(標準出力)とエラーやログ情報を分離するための一般的なプラクティスです。
*/
