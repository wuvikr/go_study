package main

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10) // 生成10以内的随机值
	fmt.Printf("send: %v\n", value)

	time.Sleep(time.Second * 5) // 等待5秒后再写入数据
	values <- value
}

func main() {
	defer close(values) // 关闭chan

	go send()
	fmt.Println("wait chanell data...")
	value := <-values // 程序会阻塞在这，直到chan中的值发送过来
	fmt.Printf("receive: %v\n", value)
	fmt.Println("main end...")
}
