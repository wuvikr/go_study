package main

import "fmt"

/*
iota是go中一个比较特殊的关键字，
可认为是一个可修改的常量，
每次换行调用后自增1，
遇到const关键字后重新置为0。
最后，同一行出现多个iota，其值不会变化。
*/
func iotaTest() {
	const (
		x1 = iota // 0
		x2 = iota // 1

		//使用 '_' 跳过2这个值
		_ // 2
		x3
		// 插入一个值
		x4 = 100
		x5
		x6 = iota // 6

		x7, x8 = iota, iota + 10
	)

	const (
		// 新的const声明语句块中，iota从0开始
		a = iota
		b = iota
	)

	fmt.Printf("x1: %v\n", x1)
	fmt.Printf("x2: %v\n", x2)
	fmt.Printf("x3: %v\n", x3)
	fmt.Printf("x4: %v\n", x4)
	fmt.Printf("x5: %v\n", x5)
	fmt.Printf("x6: %v\n", x6)
	fmt.Printf("x7: %v\n", x7)
	fmt.Printf("x8: %v\n", x8)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}

func main() {
	iotaTest()
}
