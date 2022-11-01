package main

import "fmt"

func basic() {

	// 定义一个 *string 类型的指针
	var p *string

	s := "abc"
	fmt.Printf("s: %v\n", s)

	// 使用 & 操作符生成变量 s 的内存地址
	p = &s
	fmt.Printf("p: %T\n", p)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("*p: %v\n", *p)

	// 通过指针 p 修改 s 变量的值
	*p = "xyz"
	fmt.Printf("s: %v\n", s)
	fmt.Printf("*p: %v\n", *p)
}

func arrayPointer() {
	a := [4]int{1, 2, 3, 4}
	var ptr [4]*int
	fmt.Printf("ptr: %v\n", ptr)

	for i := 0; i < len(ptr); i++ {
		ptr[i] = &a[i]
	}
	fmt.Printf("ptr: %v\n", ptr)

}

func main() {
	arrayPointer()
}
