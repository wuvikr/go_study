package main

import "fmt"

func def_slice() {
	// 声明一个字符串切片,未初始化的slice == nil
	var a []string

	// 声明一个整型切片并初始化
	var b = []int{}

	// 声明一个布尔切片并初始化
	var c = []bool{false, true}

	// 声明一个布尔切片并初始化
	//var d = []bool{false, true}

	//使用make函数声明slice并初始化
	e := make([]int, 3)

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("e: %v\n", e)

	fmt.Printf("a == nil ？ %v\n", a == nil)
	fmt.Printf("b == nil ？ %v\n", b == nil)
	fmt.Printf("c == nil ？ %v\n", c == nil)
	fmt.Printf("e == nil ？ %v\n", e == nil)

	// slice can only be compared to nil
	//fmt.Println(c == d)
}

func main() {
	def_slice()
}
