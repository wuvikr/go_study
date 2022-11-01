package main

import "fmt"

func declareArray() {
	//声明一个长度为3的数组
	var arr1 [3]int

	//声明一个长度为5，类型为int的数组并初始化赋值
	arr2 := [5]int{3, 5, 6, 6, 1}

	//声明一个长度为2，类型为string的数组并初始化赋值
	var arr3 [2]string = [2]string{"zhangsan", "lisi"}

	//声明一个长度为4，类型为bool的数组，只初始化前两个元素，未初始化元素使用默认值
	var arr4 = [4]bool{true, false}

	//声明一个类型为string的数组，自动依据初始化赋值元素个数确定数组长度
	var arr5 = [...]string{"aa", "bb", "cc"}

	//声明一个类型为bool的数组，并依据index进行赋值，未赋值的自动填充默认值
	var arr6 = [...]bool{1: true, 3: true, 5: true}

	fmt.Printf("arr1 数组类型：%T,\t\t数组元素：%v\n", arr1, arr1)
	fmt.Printf("arr2 数组类型：%T,\t\t数组元素：%v\n", arr2, arr2)
	fmt.Printf("arr3 数组类型：%T,\t数组元素：%v\n", arr3, arr3)
	fmt.Printf("arr4 数组类型：%T,\t\t数组元素：%v\n", arr4, arr4)
	fmt.Printf("arr5 数组类型：%T,\t数组元素：%v\n", arr5, arr5)
	fmt.Printf("arr6 数组类型：%T,\t数组元素：%v\n", arr6, arr6)

}

func main() {
	declareArray()
}
