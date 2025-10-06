/*
特定の型に対して、JSONへの変換ロジックを自分で定義したい場合があります。
例えば、time.TimeをデフォルトのRFC3339形式ではなく、"2006-01-02"のようなカスタムフォーマットで扱いたい場合などです。

そのためには、json.Marshalerとjson.Unmarshalerインターフェースを実装します。

・MarshalJSON() ([]byte, error): このメソッドを実装すると、json.Marshalはその型の値をこのメソッドで変換します。package main
・UnmarshalJSON(data []byte) error: このメソッドを実装すると、json.Unmarshalはこのメソッドで変換します。
*/

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// CustomTimeは"YYY-MM-DD"形式でJSON変換される
type CustomTime struct {
	time.Time
}

const layout = "2006-01-02"

// MarshalJSONはjson.Marshalerインターフェースの実装
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	// " " で囲んだバイト列を返す必要がある
	return []byte(`"` + ct.Time.Format(layout) + `"`), nil
}

// UnmarshalJSONはjson.Unmarshalerインターフェースの実装
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	// " " を取り除く
	s := string(b[1 : len(b)-1])
	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

type Event struct {
	Name string     `json:"name"`
	Date CustomTime `json:"date"`
}

func main() {
	event := Event{
		Name: "Go Conference",
		Date: CustomTime{Time: time.Now()},
	}

	// マーシャリング
	jsonData, _ := json.Marshal(event)
	fmt.Println("Marshalled:", string(jsonData))

	// アンマーシャリング
	var newEvent Event
	json.Unmarshal(jsonData, &newEvent)
	fmt.Println("Unmarshalled Name:", newEvent.Name)
	fmt.Println("Unmarshalled Date:", newEvent.Date.Time.Format("Jan 2, 2006"))
}
