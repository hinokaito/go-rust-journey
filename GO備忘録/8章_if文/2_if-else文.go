package main

import "fmt"

/*
	if文にelseを使うことで、条件がfalseの場合に別の処理を実行できます。
	if 条件 {
		処理
	} else {
		処理
	}

*/

func main() {
	scores := make(map[string]int)
	scores["saki"] = 85
	scores["yuto"] = 50

	scoring := scores["saki"]
	if scoring >= 60 {
		fmt.Println("sakiは合格です")
	} else {
		fmt.Println("sakiは不合格です")
	} // sakiは合格です

	scoring = scores["yuto"]
	if scoring >= 60 {
		fmt.Println("yutoは合格です")
	} else {
		fmt.Println("yutoは不合格です")
	} // yutoは不合格です
}
