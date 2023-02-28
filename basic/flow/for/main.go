package main

import "fmt"

func basic() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum of 1..100 is", sum)
}

/*
在go中，for循环的条件可以定义在语句块外面，注意变量的作用域区别
另外，操作变量的表达式也可以定义在语句块内
*/
func forO() {
	i := 5
	for i <= 5 {
		fmt.Printf("i: %v\n", i)
		i--
	}
}

/*
在go中，for循环提供了一种更加简便的语法糖形式，for range
*/
func forRange() {
	s := []int{1, 3, 5}

	for index, value := range s {
		fmt.Printf("index=%d value=%d\n", index, value)
	}
}

func main() {
	// basic()
	// forO()
	forRange()
}
