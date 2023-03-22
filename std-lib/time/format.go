package main

import (
	"fmt"
	"log"
	"time"
)

// go语言的时间模板比较特殊,并不是常见的 Y-m-d H:M:S
// 而是以go语言的诞生时间：2006年1月2日 15时04分 Mon Jan 为默认模板
func format() {
	now := time.Now()

	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04.05.000 Mon Jan"))

	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04.05.000 Mon Jan"))
	fmt.Println(now.Format("2006/01/02 03:04:05"))
	fmt.Println(now.Format("03:04 2006 01 02"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("03:04"))
}

func string2time() {
	now := time.Now()
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
		return
	}

	// 按照时区及指定格式解析字符串
	timeObj, err2 := time.ParseInLocation("2006/01/02 03:04:05", "2023/03/22 03:55:34", loc)
	if err2 != nil {
		log.Fatal(err2)
		return
	}
	fmt.Printf("timeObj: %v\n", timeObj)
	fmt.Printf("timeObj.Sub(now): %v\n", timeObj.Sub(now))
}

func main() {
	// format()
	string2time()
}
