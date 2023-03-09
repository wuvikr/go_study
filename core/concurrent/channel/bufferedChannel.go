package main

import (
	"fmt"
	"sync"
	"time"
)

func produce(ch chan<- int) { // 生产者一般定义为只发送类型
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
	close(ch) // 发送端负责关闭channel
	// ch <- 15 // 向已经关闭的channel中执行发送操作会引发panic
}

func consume(ch <-chan int) { // 消费者一般定义为只接收类型
	for n := range ch {
		fmt.Printf("n: %v\n", n)
	}
}

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()

	wg.Wait()
}
