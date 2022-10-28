package main

import "fmt"

func sum(a int, b int) (c int) {
	c = a + b
	return c
}

func updateAgeTo18(age int) {
	age = 18
}

func updateAgeTo19(age *int) {
	*age = 19
}

func sliceUpdate(s1 []int) {
	s1[0] = 1
}

func main() {
	/* i := sum(1, 4)
	fmt.Printf("i: %v\n", i) */

	/* age := 27
	updateAgeTo18(age)
	fmt.Printf("age: %v\n", age) */

	/* age := 27
	updateAgeTo19(&age)
	fmt.Printf("age: %v\n", age) */

	slice1 := []int{4, 5, 6}
	sliceUpdate(slice1)
	fmt.Printf("slice1: %v\n", slice1)

}
