package main

import (
	"fmt"
	"time"
)

type signal struct{}

func worker() {
	fmt.Println("worker is working...")
	time.Sleep(time.Second * 5)
}

// 等调用函数执行完，返回signal作为Goroutine退出的通知信号
func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		fmt.Println("worker start to work...")
		f()
		c <- signal{}
	}()
	return c
}

func main() {
	fmt.Println("start a worker...")
	c := spawn(worker)
	<-c // 直到channel中信号发送过来前，main Goroutine会一直被阻塞
	fmt.Println("worker work done!")
}
