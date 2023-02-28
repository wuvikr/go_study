package main

import "fmt"

func anonymousStruct() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "哒哒"
	user.Age = 18
	fmt.Printf("%#v\n", user)
}

// Person 结构体Person类型
type Person struct {
	string
	int
}

func main() {
	//anonymousStruct()

	p1 := Person{
		"丁丁",
		12,
	}
	//fmt.Printf("p1: %#v\n", p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
