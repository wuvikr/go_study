package main

import "fmt"

/*
iota是go中一个比较特殊的关键字，
可认为是一个可修改的常量，
每次调用后自增1，
遇到const关键字后重新置为0。
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
	)

	fmt.Printf("x1: %v\n", x1)
	fmt.Printf("x2: %v\n", x2)
	fmt.Printf("x3: %v\n", x3)
	fmt.Printf("x4: %v\n", x4)
	fmt.Printf("x5: %v\n", x5)
	fmt.Printf("x6: %v\n", x6)

}

func main() {
	//未使用的常量不会报错
	const pi = 3.1415926

	//常量不可重新赋值，会报错
	//pi=3.15

	const (
		a = 1
		b = 2
		// 一次性给多个常量赋值时，未赋值的常量会继承上一个赋值语句中的值
		c
	)

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)

	iotaTest()
}
