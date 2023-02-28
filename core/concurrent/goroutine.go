package main

import (
	"fmt"
	"time"
)

func printMsg(s string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100) // sleep 100ms
	}
}

func main() {
	go printMsg("golang") // 启用一个新的协程
	go printMsg("python")

	time.Sleep(time.Millisecond * 100)
	fmt.Println("main end...")
}
