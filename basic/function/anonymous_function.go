package main

import "fmt"

func main() {
	// 将匿名函数保存到变量中
	add := func(x, y int) {
		fmt.Println(x + y)
	}

	// 通过变量调用匿名函数
	add(10, 20)

	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
