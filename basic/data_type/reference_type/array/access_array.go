package main

import "fmt"

func main() {
	var a = [4]int{3, 6, 9, 5}

	// 给下标为1的元素重新赋值
	a[1] = 10
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a[1]: %v\n", a[1])

	// 访问数组最后一个元素，先用函数得到数组长度，长度减1即为最后一个元素的下标
	b := len(a) - 1
	fmt.Printf("b: %v\n", b)
	fmt.Printf("a[b]: %v\n", a[b])
}
