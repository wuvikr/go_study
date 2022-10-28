package main

import "fmt"

func foo() {

	//if/else
	x := 1
	if n := "abc"; x > 0 {
		println(n[2])
	} else if x < 0 {
		println(n[1])
	} else {
		println(n[0])
	}

	// nested if
	var a int = 100
	var b int = 200
	if a == 100 {
		if b == 200 {
			fmt.Printf("a 的值为 100， b 的值为 200。\n")
		}
	}
}

func localVar() {
	if a, b, c := 9, 5, 9; (a+b)*c >= 100 {
		fmt.Println("YES")
	} else {
		fmt.Println("No")
	}
}

func main() {

	foo()

	localVar()

	x := 14
	if x%2 == 0 {
		fmt.Println("x是一个偶数！")
	}
}
