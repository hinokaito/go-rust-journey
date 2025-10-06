package main

import (
	"fmt"
	"time"
)

/*時間の計算
・time.Add(d Duration): time.TimeにDurationを加算し、新しいtime.Timeを返します
・time.Sub(t2 Time): 2つのtime.Timeの差をDurationとして返します。
・time.AddDate(years, months, days int): 年月日単位で加算・減算を行います
*/

func main() {
	now := time.Now()

	// 10分後の時刻
	tenMinutesLater := now.Add(10 * time.Minute)
	fmt.Println("10分後:", tenMinutesLater)

	// 1時間30分前の時刻
	oneHour30MinAgo := now.Add(-1*time.Hour - 30*time.Minute)
	fmt.Println("1時間30分前:", oneHour30MinAgo)

	// 1年2ヶ月3日後の日付
	futureDate := now.AddDate(1, 2, 3)
	fmt.Println("1年2ヶ月3日後:", futureDate)

	// 2つの時刻の差
	diff := tenMinutesLater.Sub(now)
	fmt.Println("差は %.0f 秒です\n", diff.Seconds())
}
