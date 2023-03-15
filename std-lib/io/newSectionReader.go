package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	// 截取指定部分
	sr := io.NewSectionReader(r, 5, 17)

	if _, err := io.Copy(os.Stdout, sr); err != nil {
		log.Fatal(err)
	}
}
