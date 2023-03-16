package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("ABC DEF GHI JKL")
	r2 := bufio.NewReader(r)

	s1, _ := r2.ReadSlice(' ')
	fmt.Printf("s1: %q\n", s1)

	s2, _ := r2.ReadSlice(' ')
	fmt.Printf("s2: %q\n", s2)

	s3, _ := r2.ReadSlice(' ')
	fmt.Printf("s3: %q\n", s3)

	s4, _ := r2.ReadSlice(' ')
	fmt.Printf("s4: %q\n", s4)
}
