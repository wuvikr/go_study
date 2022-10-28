package main

import (
	"fmt"
)

func basic() {
	i := 3

	switch i {
	case 0:
		fmt.Print("Zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")
	case 3:
		fmt.Print("three...")
	default:
		fmt.Print("no match...")
	}

	fmt.Println("OK")
}

/*
go中，switch case，switch可以省略判断，而将表达式放在case中。
*/
func caseIF() {
	age := 19
	switch {
	case age > 18:
		fmt.Println("成年")
	case age < 18:
		fmt.Println("未成年")
	}
}

// case支持匹配多个表达式
func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "北京", "上海", "深圳", "广州":
		region, continent = "中国", "亚洲"
	case "东京", "大阪", "名古屋", "横滨":
		region, continent = "日本", "亚洲"
	case "纽约", "洛杉矶", "旧金山":
		region, continent = "美国", "北美洲"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

/*
go的switch case不需要break关键字，
匹配到一个后就会自动退出
如果想继续执行后么的case可以使用fallthrough关键字
注意：使用fallthrough关键字后，不需要case能匹配上，会直接执行下一个case中的语句块
*/
func caseFT() {
	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	// 这里15明显不可能大于100，但由于fallthrough关键字的特性，会直接执行下一个case中的语句块
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
	case num < 200:
		fmt.Printf("%d is less than 200", num)
	}
}

func main() {
	basic()
	caseIF()
	city := "上海"
	s, s2 := location(city)
	fmt.Printf("%s 位于 %s%s\n", city, s2, s)

	caseFT()
}
