package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("hello world")

	f, _ := os.Open("README.md")
	r2 := bufio.NewReader(r)
	s, _ := r2.ReadString(' ')
	fmt.Printf("s: %v\n", s)

	r2.Reset(f) // 丢弃缓冲中的数据，清除错误
	s2, _ := r2.ReadString(' ')
	fmt.Printf("s2: %v\n", s2)
}
