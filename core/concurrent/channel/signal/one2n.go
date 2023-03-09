package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done\n", i)
}

type signal struct{}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-groupSignal
			fmt.Printf("worker %d: start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	// 这里的wg.Wait()为啥要单独启一个Goroutine?
	// wg.Wait()会阻塞所有的子Goroutine,如果这里不放在一个新的Goroutine中会阻塞main Goroutine
	// 导致main函数中，调用spawnGroup函数下面的语句都无法执行，导致deadlock
	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

// main goroutine 创建了一组 5 个 worker goroutine，
// 这些 Goroutine 启动后会阻塞在名为 groupSignal 的无缓冲 channel 上。
// main goroutine 通过close(groupSignal)向所有 worker goroutine 广播“开始工作”的信号，
// 收到 groupSignal 后，所有 worker goroutine 会“同时”开始工作。
func main() {
	fmt.Println("start a group of workers...")
	groupSignal := make(chan signal)
	c := spawnGroup(worker, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	close(groupSignal)
	<-c
	fmt.Println("the group of workers work done!")
}
