package main

import "fmt"

// 斐波拉契数列
func Fibonacci(i int) int {

	if i == 1 || i == 2 {
		return 1
	}

	return Fibonacci(i-1) + Fibonacci(i-2)
}

// 阶乘
func factorial(i int) int {
	if i == 1 {
		return 1
	}

	return factorial(i-1) * i
}

func main() {
	/* f4 := Fibonacci(4)
	fmt.Printf("f4: %v\n", f4) */

	fmt.Println(factorial(5))
}
