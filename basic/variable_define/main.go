package main

import "fmt"

// 匿名变量，函数可以只写返回值类型，不写名称
func foo() (int, string) {
	return 1, "hi"
}

func main() {
	var name string = "wuvikr"

	// 短变量声明，只能用在函数内部，自动推断类型。
	age := 18

	// 批量声明变量
	var (
		id, rank = 1, "high"
		sex      string
		num      int64 = 13126274829
	)

	// 只需要返回值中的一个，另一个不需要的返回值可以用_接收
	x, _ := foo()
	_, y := foo()

	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("id: %v\n", id)
	fmt.Printf("rank: %v\n", rank)
	fmt.Printf("sex: %v\n", sex)
	fmt.Printf("num: %v\n", num)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
}
