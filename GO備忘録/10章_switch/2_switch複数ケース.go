package main

import "fmt"

func main() {
	day := "Saturday"

	switch day {
	case "Saturday", "Sunday": // 土曜日と日曜日を 「or」 としてまとめて処理できます
		fmt.Println("週末です")
	default:
		fmt.Println("平日です")
	}
	// 土日で二つ分けるのも良いですが、このほうが視認性的にもコーディング量的にもgoodです
}
