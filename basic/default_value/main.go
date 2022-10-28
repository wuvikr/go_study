package main

import "fmt"

func default_value() {
	var a int
	var b byte
	var f float32
	var t bool
	var s string
	var r rune
	var arr []int

	fmt.Printf("default value of int %d\n", a)
	fmt.Printf("default value of byte %d\n", b)
	fmt.Printf("default value of float32 %f\n", f)
	fmt.Printf("default value of bool %t\n", t)
	fmt.Printf("default value of string %s\n", s)
	fmt.Printf("default value of rune %d\n", r)
	fmt.Printf("default value of arr %v\n", arr)
}

func main() {
	default_value()
}
