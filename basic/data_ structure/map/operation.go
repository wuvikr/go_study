package main

import "fmt"

func main() {
	// mapAdd()
	// mapAccess()
	// ifKeyExist()
	// mapDelete()
	mapIterate()
}

func mapAdd() {
	//注意：make创建的map已经是初始化的了，如果仅仅声明一个map，使用下面的方式会报错
	//	var studentScore map[string]int 会报错
	studentScore := make(map[string]int, 8)
	studentScore["张三"] = 90
	studentScore["小明"] = 100

	fmt.Printf("studentScore: %v\n", studentScore)
	fmt.Printf("studentScore[\"小明\"]: %v\n", studentScore["小明"])
	fmt.Printf("type of studentScore : %T\n", studentScore)
}

func mapAccess() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	fmt.Println("Bob's age is", studentsAge["bob"])
	//注意：在go中，使用下标法访问map时，即使访问一个不存在的键，go也不会执行panic，而是会返回类型的默认值。
	fmt.Println("Christy's age is", studentsAge["christy"])
}

/*
go中，map的下标表示法会返回两个值，第一个是下标所对应的值，第二个是布尔型值，用来表示key是否存在。
因此，当需要判断某个key是否存在时，不能直接以某个key是否有返回值来作为判断。
*/
func ifKeyExist() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}
}

func mapDelete() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 77
	scoreMap["李四"] = 88
	scoreMap["王五"] = 99
	fmt.Printf("scoreMap: %v\n", scoreMap)

	// 将王五从map中删除
	delete(scoreMap, "王五")
	fmt.Printf("scoreMap: %v\n", scoreMap)

	// 删除一个不存在的key刘六,go并不会panic
	delete(scoreMap, "刘六")
	fmt.Printf("scoreMap: %v\n", scoreMap)
}

func mapIterate() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	for name, age := range studentsAge {
		fmt.Printf("%s \t%d\n", name, age)
	}

	//for range只有一个返回值的时候，返回key
	for value := range studentsAge {
		fmt.Printf("%s\n", value)
	}
}
