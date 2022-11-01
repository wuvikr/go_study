package main

import "fmt"

/*
adder()函数返回一个累加函数，
这个函数作用域中包含有一个外部变量sum，
sum的值会随着函数的每次调用而更新。
*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func startingValueAdder(sum int) func(int) int {
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	// 这里初始化的pos后的变量可以理解成生成了下面这样一个函数域
	pos := adder()
	/*
		var sum int = 0
		func sums(x int) int{
			sum +=x
			return sum
		}
	*/

	// 每次调用pos()函数，函数域内部的sum值都会被存储和更新
	fmt.Printf("pos(5): %v\n", pos(5))
	fmt.Printf("pos(5): %v\n", pos(5))
	fmt.Printf("pos(5): %v\n", pos(5))

	// 重新对pos变量进行初始化后，sum的值也会被初始化
	pos = adder()
	fmt.Printf("pos(5): %v\n", pos(5))

	f := startingValueAdder(10)
	fmt.Printf("f(5): %v\n", f(5))
	fmt.Printf("f(5): %v\n", f(5))
	fmt.Printf("f(5): %v\n", f(5))
}
