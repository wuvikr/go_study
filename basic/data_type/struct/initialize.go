package main

import "fmt"

type person struct {
	name      string
	age       int
	sex, city string
	adult     bool
}

/*
结构体中的字段可以使用"."来访问，
因此，当然也可以用来赋值初始化
*/
func structInitializeWay01() {
	var wanger person

	wanger.name = "wanger"
	wanger.age = 22
	wanger.sex = "female"
	wanger.city = "shanghai"
	wanger.adult = true

	fmt.Printf("wanger: %#v\n", wanger)
}

/*
声明结构体变量的同时进行赋值，
使用“字段名： 值”的格式，
这种方式的初始化因为指定了字段名，所以可以不用按结构体定义的顺序来初始化，
注意：和其他语言不太一样，最后一行也需要加上 “,”。
*/
func structInitializeWay02() {
	lisi := person{
		age:  15,
		name: "lisi",
		city: "BeiJing",
		sex:  "male",
	}

	fmt.Printf("lisi: %#v\n", lisi)
}

/*
初始化结构体时，具体字段可以省略不写，
但这种初始化方式有以下几点需要注意：

1. 赋值的顺序会和结构体中的字段声明顺序一致；
2. 必须初始化结构体中的所有字段
3. 这种方式不能和“字段名：值”的方式混用
*/
func structInitializeWay03() {
	zhangsan := person{
		"zhangsan",
		33,
		"male",
		"Shanghai",
		true,
	}
	fmt.Printf("zhangsan: %#v\n", zhangsan)
}

func main() {
	structInitializeWay01()
	structInitializeWay02()
	structInitializeWay03()
}
