package main

import "fmt"

func sliceCap() {

	var numbers []int
	for i := 0; i < 10; i++ {
		// 使用append函数向numbers切片中追加元素，注意append不支持数组，数组是不可变长度。
		numbers = append(numbers, i)

		// 请注意观察切片的容量cap的变化
		fmt.Printf("i=%d\tlen=%d\tcap=%d\t%v\n", i, len(numbers), cap(numbers), numbers)
	}
}

func sliceAppend() {

	var a []int
	fmt.Printf("a: %v\n", a)

	// 追加单个元素
	a = append(a, 1)
	fmt.Printf("a: %v\n", a)

	// 追加多个元素
	a = append(a, 2, 3, 4)
	fmt.Printf("a: %v\n", a)

	// 追其他切片
	b := []int{5, 6}
	a = append(a, b...)
	fmt.Printf("a: %v\n", a)
}

/*
go中没有从切片中删除元素的内置函数。
可使用切片运算符 s[i:p] 来新建一个仅包含所需元素的切片。
*/
func sliceDelete() {
	s1 := []int{1, 2, 3, 4, 5}

	// 删除索引为2的元素
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("s1: %v\n", s1)
}

func main() {
	sliceAppend()
	//sliceCap()
	//sliceDelete()
}
