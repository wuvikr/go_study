package main

import "fmt"

/*
切片表达式s[i:p]:
s表示数组或切片

● i表示指向要添加到新切片的数组或切片的第一个元素的指针。

● p表示可用于新切片的数组或切片中的最后一个元素的位置。

简单理解就是切片表达式中的i和p表示一个索引范围 （左包含，右不包含） [i]中索引位置i处的元素。
*/
func main() {
	array := [7]int{4, 2, 5, 7, 9, 8, 1}

	slice1 := array[1:6]
	fmt.Printf("slice1:%v\n", slice1)
	fmt.Printf("len(slice1):%v\n", len(slice1))
	fmt.Printf("cap(slice1):%v\n\n", cap(slice1))

	// 从下标3到末尾
	slice2 := slice1[3:]
	fmt.Printf("slice2:%v\n", slice2)
	fmt.Printf("len(slice2):%v\n", len(slice2))
	fmt.Printf("cap(slice2):%v\n\n", cap(slice2))

	// 切片拷贝
	slice3 := slice2[:]
	fmt.Printf("slice3:%v\n", slice3)
	fmt.Printf("len(slice3):%v\n", len(slice3))
	fmt.Printf("cap(slice3):%v\n", cap(slice3))

}
