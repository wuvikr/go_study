package main

import "fmt"

type Dog struct {
	name string
	age  int
}

func (dog Dog) eat() {
	fmt.Println("狗吃饭！")
}

func main() {
	dog1 := Dog{
		name: "xiaohei",
		age:  3,
	}

	dog1.eat()
}
