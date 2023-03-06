package main

import (
	"fmt"
	"strconv"
)

func plusOne(digits []int) []int {
	var s string
	var slice1 = []int{}
	for _, v := range digits {
		fmt.Println(v)
		s = s + strconv.Itoa(v)
	}
	i, _ := strconv.Atoi(s)
	i++
	s = strconv.Itoa(i)
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		a, _ := strconv.Atoi(string(s[i]))
		slice1 = append(slice1, a)
	}
	return slice1
}

func main() {
	var s = []int{9, 9}
	fmt.Println(plusOne(s))
}
