package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取所有环境变量
	s := os.Environ()
	fmt.Printf("s: %v\n", s)

	// 设置环境变量
	os.Setenv("ping", "pong")

	// 获取指定环境变量
	s2 := os.Getenv("ping")
	fmt.Printf("s2: %v\n", s2)

	// 查找某个环境变量是否存在
	s3, b := os.LookupEnv("ping")
	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("b: %v\n", b)

	// 字符串中引用环境变量
	fmt.Println(os.ExpandEnv("My gopath is ${GOPATH}"))

	// 清空环境变量
	// os.Clearenv()
}
