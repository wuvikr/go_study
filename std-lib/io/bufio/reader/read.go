package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func read() {
	r := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	r2 := bufio.NewReader(r)

	p := make([]byte, 10)

	for {
		n, err := r2.Read(p) // 根据指定的数组长度读取buff，并将读到的个数返回
		if err == io.EOF {
			break
		} else {
			fmt.Printf("string(p[0:n]): %v\n", string(p[0:n]))
		}
	}
}

func main() {
	read()
}
