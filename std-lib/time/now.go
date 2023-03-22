package main

import (
	"fmt"
	"time"
)

// 当前时间
func now() {
	now := time.Now()
	fmt.Printf("the type of now is : %T, current time: %v\n", now, now)
	// current time: 2023-03-22 14:28:29.2827255 +0800 +08 m=+0.001998001

	year := now.Year()
	fmt.Printf("year: %v\n", year)

	month := now.Month()
	fmt.Printf("month: %v\n", month)

	day := now.Day()
	fmt.Printf("day: %v\n", day)

	hour := now.Hour()
	fmt.Printf("hour: %v\n", hour)

	minute := now.Minute()
	fmt.Printf("minute: %v\n", minute)

	second := now.Second()
	fmt.Printf("second: %v\n", second)

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T %T %T %T %T %T\n", year, month, day, hour, minute, second)

}

func timeStamp() {
	ts := time.Now().Unix()
	fmt.Printf("the type of timeStamp is: %T, value: %v\n", ts, ts)

	// 时间戳转化为普通时间格式
	timeObj := time.Unix(ts, 0)
	fmt.Printf("timeObj: %v\n", timeObj)

	i := timeObj.Year()
	m := timeObj.Month()
	i2 := timeObj.Day()
	i3 := timeObj.Hour()
	i4 := timeObj.Minute()
	i5 := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", i, m, i2, i3, i4, i5)

	year, m2, day := timeObj.Date()
	fmt.Printf("year : %d, month : %d, day : %d ", year, m2, day)
}

func main() {
	// now()
	timeStamp()
}
