package main

import "fmt"

type person struct {
	name string
	age  int
	sex  int
}

// 对于person结构体来说，其中一些字段的默认值是有实际意义的
// 比如年龄，因此想用-1来表示年龄未知
// 还比如sex字段，有时会用0来表示女性，1表示男性，这样sex的默认值也是有意义的
func newDefaultPerson() *person {
	return &person{
		name: "未知",
		age:  -1,
		sex:  2,
	}
}

func newPerson(name string, age, sex int) *person {
	return &person{
		name: name,
		age:  age,
		sex:  sex,
	}
}

func main() {
	p1 := person{}
	fmt.Printf("p1: %v\n", p1)

	p2 := new(person) // 返回指针类型
	fmt.Printf("p2: %v\n", p2)

	p3 := newDefaultPerson()
	fmt.Printf("p3: %v\n", p3)

	p4 := newPerson("abby", 18, 0)
	fmt.Printf("p4: %v\n", p4)
}
