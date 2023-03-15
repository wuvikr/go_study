package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	b := make([]byte, 8)

	// 带缓冲区的copy
	if _, err := io.CopyBuffer(os.Stdout, r1, b); err != nil {
		log.Fatal(err)
	}
	if _, err := io.CopyBuffer(os.Stdout, r2, b); err != nil {
		log.Fatal(err)
	}
}
