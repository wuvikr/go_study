package main

import "fmt"

// type funcT func(int, int) int

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	// var f1 funcT
	// f1 = max
	// fmt.Printf("f1(3, 5): %v\n", f1(3, 5))

	f2 := min
	fmt.Printf("f2(3, 5): %v\n", f2(3, 5))
}
