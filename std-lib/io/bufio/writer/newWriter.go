package main

import (
	"bufio"
	"os"
)

func main() {
	f, _ := os.OpenFile("README.md", os.O_RDWR, 0644)
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("hello golang...")
	w.Flush() // 强制发送缓冲区数据到io.writer
}
