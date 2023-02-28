package main

import "fmt"

func basic() {
	sum := 0
	for num := 1; num <= 100; num++ {
		if num%5 == 0 {
			continue
		}
		sum += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sum)
}

func labelContinue() {
	var arr2 = [][]int{
		{1, 4, 6, 7, 9},
		{2, 5, 7, 9, 4},
		{3, 5, 1, 6, 7},
	}

outerloop:
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			if arr2[i][j] == 7 {
				fmt.Printf("found 7 at [%d, %d]\n", i, j)
				continue outerloop
			}
		}
	}
}

func labelBreak() {
	var arr2 = [][]int{
		{1, 4, 6, 7, 9},
		{2, 5, 3, 9, 4},
		{0, 5, 1, 6, 7},
	}

outerloop:
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			if arr2[i][j] == 3 {
				fmt.Printf("found 7 at [%d, %d]\n", i, j)
				break outerloop
			}
		}
	}
}

func main() {
	// basic()
	// labelContinue()
	labelBreak()
}
