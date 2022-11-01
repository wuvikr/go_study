package main

import "fmt"

func forloop(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func rangeLoop(arr [5]int) {
	for index, value := range arr {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}
}

func main() {
	var arr1 = [...]int{3, 3, 5, 6, 8}
	//forloop(arr1)
	rangeLoop(arr1)
}
