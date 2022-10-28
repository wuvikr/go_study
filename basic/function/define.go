package main

import "fmt"

// 最终返回值以return关键字返回内容为准
func max(a int, b int) (max int) {
	c := a + b
	return c
}

// 多个返回值
func HighOrLow(a, b int) (high, low int) {
	if a == b {
		panic("a == b!")
	}
	if a > b {
		high = a
		low = b
	} else {
		high = b
		low = a
	}
	return high, low
}

func printMeg(name, city, sex string, age int) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("sex: %v\n", sex)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("city: %v\n", city)
}

// return后面可以不写，默认返回定义的形式变量
func testReturn() (name string, age int) {
	name = "张三"
	age = 77
	age = 00
	return
}

// 可变长参数
func sum(agrs ...int) (sum int) {
	sum = 0
	for _, v := range agrs {
		sum += v
	}
	return sum
}

func main() {
	/* v := max(1, 3)
	fmt.Printf("v: %v\n", v) */

	/* higt, low := HighOrLow(2, 3)
	fmt.Printf("higt: %v\n", higt)
	fmt.Printf("low: %v\n", low) */

	/* s, i := testReturn()
	fmt.Printf("s: %v\n", s)
	fmt.Printf("i: %v\n", i) */

	sum(1, 2, 3, 4, 5)
}
