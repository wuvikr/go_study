package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var lock sync.Mutex

func add(i int) {
	defer wg.Done()
	lock.Lock()
	time.Sleep(time.Millisecond * 10)
	i += 1
	fmt.Printf("i++: %v\n", i)
	lock.Unlock()
}

func sub(i int) {
	defer wg.Done()
	lock.Lock()
	time.Sleep(time.Millisecond * 2)
	i -= 1
	fmt.Printf("i--: %v\n", i)
	lock.Unlock()
}

func main() {
	var a int = 100
	for i := 0; i <= 100; i++ {
		go add(a)
		wg.Add(1)
		go sub(a)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Printf("end... a =  %v\n", a)
}
