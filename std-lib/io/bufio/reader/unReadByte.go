package main

import (
	"bufio"
	"fmt"
	"strings"
)

func unReadByte() {
	r := strings.NewReader("ABCDEFG")
	r2 := bufio.NewReader(r)

	b, _ := r2.ReadByte() // 从index为0开始读取并返回一个字节，并将下次要读取字节的index往后移动一位，没有数据可读会报错。
	fmt.Printf("b: %c\n", b)

	c, _ := r2.ReadByte()
	fmt.Printf("c: %c\n", c)

	r2.UnreadByte() // 将当前读取字节的index往回移动一位
	d, _ := r2.ReadByte()
	fmt.Printf("d: %c\n", d)
}

func main() {
	unReadByte()
}
