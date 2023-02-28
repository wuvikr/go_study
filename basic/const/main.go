package main

import "fmt"

func main() {
	//未使用的常量不会报错
	const pi = 3.1415926

	//常量不可重新赋值，会报错
	//pi=3.15

	const (
		a    = 1
		b, c = 2, 4
		// 使用const语句块一次给多个常量赋值时，未赋值的常量会继承上一个赋值语句中的值
		d, e
	)

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)
}
