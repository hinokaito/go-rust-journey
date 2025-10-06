/*不定なJSONの扱い
APIｋらのレスポンスなど、JSONの構造が固まっていない、あるいは動的に変わる場合があります。そのような場合はmap[string]interface{}を使います。

interface{}(Go 1.18以降はanyのエイリアス)は、あらゆる型を格納できる特別な型です。
これにより、どのような構造のJSONでもマップとして受け取ることができます。
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	jsonStr := `{"id":99, "name":"Unknown", "details":{"age":30, "hobbies":["reading","sports"]}}`
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Full map:", data)

	// 値にアクセスするには型アサーションが必要
	id := data["id"].(float64) // JSONの数値はデフォルトでfloat64になる
	name := data["name"].(string)
	fmt.Printf("ID: %.0f, Name: %s\n", id, name)

	details := data["details"].(map[string]interface{})
	age := details["age"].(float64)
	fmt.Printf("Age: %.0f\n", age)
}

/*
この方法は柔軟ですが、値を取り出す際に型アサーションが必要になり、型安全性が失われるため、コードが頬雑になりがちです。
構造がわかっている場合は、極力構造体を定義するべきです。
*/
