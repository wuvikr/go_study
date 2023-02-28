package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	// goroutine执行完登记-1
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		go hello(i)
		// 每启动一个goroutine就登记+1
		wg.Add(1)
	}

	// 等待所有登记的goroutine执行结束
	wg.Wait()
	fmt.Println("func main is end...")
}
