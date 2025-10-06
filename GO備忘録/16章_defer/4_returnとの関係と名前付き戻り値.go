package main

import "fmt"

/*
defer文は、return文が実行された後、そして関数が実際に呼び出し元に制御を返す前に実行されます。
この性質を利用すると、名前付き戻り値(Named Return Values)の値をdefer内で変更することも可能です。
*/

// 戻り値にiという名前を付けている
func increment() (i int) {
	defer func() {
		i++ // 戻り値iをインクリメント
	}()
	return 1 // ここでiは1になるが...
}

func main() {
	fmt.Println(increment()) // 2
}

/*処理の流れ:
1. increment関数が呼び出される。
2. deferで無名関数が登録される。
3. return 1が実行される。この時点で、戻り値iに1がセットされる。
4. 関数が終了する直前に、deferで登録されていた無名関数func() {i++}が実行される
5. 戻り値iがインクリメントされ、1から2になる。
6. increment関数が2を返す
*/
