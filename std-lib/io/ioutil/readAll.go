package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("gogogo!")

	b, err := ioutil.ReadAll(r)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("b: %s\n", b)
}
