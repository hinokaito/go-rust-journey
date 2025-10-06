package main

import "fmt"

/*
	if文にif-elseを使うことで、条件がfalseの場合に別の処理を実行できます。

	if 条件 {
		処理
	} else if 条件 {
		処理
	} else {
		処理
	}

	複数else ifを使うこともでき、if 条件{処理}else if 条件{処理} else if 条件{処理} else{処理}というように書くこともできます。
*/

func main() {
	scores := make(map[string]int)
	scores["saki"] = 85
	scores["yuto"] = 50

	scoring := scores["saki"]
	if scoring >= 85 {
		fmt.Println("sakiは優秀です")
	} else if scoring >= 60 {
		fmt.Println("sakiは合格です")
	} else if scoring >= 50 {
		fmt.Println("sakiは不合格ですが、再テストを受けられます")
	} else {
		fmt.Println("sakiは不合格です")
	} // sakiは優秀です

	scoring = scores["yuto"]
	if scoring >= 90 {
		fmt.Println("yutoは優秀です")
	} else if scoring >= 60 {
		fmt.Println("yutoは合格です")
	} else if scoring >= 50 {
		fmt.Println("yutoは不合格ですが、再テストを受けられます")
	} else {
		fmt.Println("yutoは不合格です")
	} // yutoは不合格ですが、再テストを受けられます
}
