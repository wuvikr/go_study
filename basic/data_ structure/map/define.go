package main

import "fmt"

func def_map() {
	// 声明 map(未初始化)
	var studentId map[string]int
	fmt.Printf("studentId: %v\n", studentId)

	// 使用make函数声明并初始化一个空的map
	studentSex := make(map[string]string)
	fmt.Printf("studentSex: %v\n", studentSex)

	// 声明map并赋值
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}
	fmt.Printf("studentsAge: %v\n", studentsAge)
}

func main() {
	def_map()
}
