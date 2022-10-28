package main

import "fmt"

/*
go中，切片直接赋值或者切片表达式产生的新切片，默认都是地址传递，
即原切片会和新切片共享一块内存地址，
修改其中任何一个都会影响到另一个
*/
func shallowCope() {
	s1 := make([]int, 3)
	fmt.Printf("s1: %v\n", s1)

	s2 := s1
	fmt.Printf("s2: %v\n", s2)
	s2[0] = 100

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)

}

/*
go提供了内置函数copy()，可以将一个切片的数据复制到另外一个切片空间中。
copy() 函数的使用格式如下：
copy(dst, src []Type)
*/
func deepCope() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 5, 5)
	s3 := make([]int, 3, 3)
	s4 := make([]int, 8, 8)

	// copy() 函数的目标切片需要提前声明初始化
	copy(s2, s1)
	copy(s3, s1)
	copy(s4, s1)
	fmt.Printf("s1切片：%v\n", s1)
	fmt.Printf("s2切片：%v\n", s2)
	fmt.Printf("s3切片：%v\n", s3)
	fmt.Printf("s4切片：%v\n", s4)
	fmt.Println()

	s3[0] = 100
	fmt.Printf("s1切片：%v\n", s1)
	fmt.Printf("s2切片：%v\n", s2)
	fmt.Printf("s3切片：%v\n", s3)
	fmt.Printf("s4切片：%v\n", s4)
}

func main() {
	// shallowCope()
	deepCope()
}
