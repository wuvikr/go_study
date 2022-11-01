package main

import "fmt"

func main() {
	var i *int
	*i = 1

	// i := new(int)
	// *i = 1

	fmt.Printf("*i : %v", *i)
}
