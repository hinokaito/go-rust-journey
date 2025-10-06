package main

import "fmt"

/*
if文は、条件がtrueの場合に、その条件に対応する処理を実行します。
if 条件 {
	処理
}


*/

func main() {

	scores := make(map[string]int)
	scores["saki"] = 85
	scores["yuto"] = 50

	// saki の点数を採点する
	scoring := scores["saki"]

	// もし saki の点数が60以上ならば
	if scoring >= 60 {
		fmt.Println("sakiは合格です")
	} // sakiは合格です

	// yuto の点数を採点する
	scoring = scores["yuto"]

	// もし yuto の点数が60以上ならば
	if scoring >= 60 {
		fmt.Println("yutoは合格です")
	} // 何も表示されない
}
