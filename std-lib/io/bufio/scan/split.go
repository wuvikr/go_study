package main

import (
	"bufio"
	"fmt"
	"strings"
)

func splitWords() {
	r := strings.NewReader("I love you.")
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		fmt.Printf("s.Text(): %v\n", s.Text())
	}
}

func main() {
	splitWords()
}
