/*構造体フィールドタグ
JSONのキー名をGoのフィールド名と変えたい場合や、特定の条件下でフィールドを無視したい場合にフィールドタグを使います。バッククォートで囲んで指定します。

json:"<キー名>,<オプション>"の形式で記述します。

・キー名の変更: JSONではsnake_caseが一般的なため、GoのPascalCaseから変更するのによく使われます。

・omitempty: フィールの値がゼロ値(数値の0, false, "", nilスライスなど)の場合、JSONの出力からそのフィールドを除外します。

・-(ハイフン): このフィールドを常にJSONの変換対象から除外します。
*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID          int    `json:"user_id"`
	Name        string `json:"user_name"`
	Email       string `json:"user_email"` // emailが空文字列なら出力しない。
	Password    string `json:"-"`          // パスワードは常に出力しない。
	privateNote string // 非公開フィールドは元々変換されない
}

func main() {
	user1 := User{
		ID:       101,
		Name:     "Ken Tanaka",
		Email:    "ken@example.com",
		Password: "password123",
	}
	user2 := User{
		ID:       102,
		Name:     "Taro Yamada",
		Password: "password456",
	}

	json1, _ := json.Marshal(user1)
	json2, _ := json.Marshal(user2)

	fmt.Println("User 1:", string(json1))
	fmt.Println("User 2:", string(json2))
}

/*
User 1: {"user_id":101,"user_name":"Ken Tanaka","user_email":"ken@example.com"}
User 2: {"user_id":102,"user_name":"Taro Yamada","user_email":""}
*/
