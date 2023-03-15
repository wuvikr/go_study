package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("hello world!")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
