package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	b, _ := ioutil.ReadFile("./README.md")
	fmt.Printf("string(b): %v\n", string(b))
}

func main() {
	readFile()
}
