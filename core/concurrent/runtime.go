package main

import (
	"fmt"
	"runtime"
)

func showMeg01(str string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("showMeg: %v\n", str)
	}
}

func showMeg02() {
	for i := 0; i < 10; i++ {
		if i >= 5 {
			runtime.Goexit()
		}
		fmt.Println(i)
	}
}

func main() {
	// go showMeg01("golang")
	// runtime.Gosched() // 当该进程被分配到时间片可以执行时，主动将时间片让出给其他协程

	// go showMeg02()

	// fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU()) // 查看本机最大CPU核心数
	// runtime.GOMAXPROCS(1)                                  // 设置协程使用的最大CPU核心数
	fmt.Println("main end...")
}
