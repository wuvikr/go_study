package main

import "fmt"

func basic() {
	var a int = 1
	var p1 *int = &a
	fmt.Printf("p1: %v\n", *p1)

	var b int = 10
	var p2 *int = &b
	fmt.Printf("p2: %v\n", *p2)

	var pp **int = &p1
	fmt.Printf("pp: %v\n", **pp)
	pp = &p2
	fmt.Printf("pp: %v\n", **pp)
}

func foo(pp **int) {
	var b int = 20
	var p1 *int = &b
	(*pp) = p1
}

func main() {
	// basic()

	var a int = 2
	var p *int = &a
	fmt.Printf("p: %v\n", *p)
	foo(&p)
	fmt.Printf("p: %v\n", *p)
}
