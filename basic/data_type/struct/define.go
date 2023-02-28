package main

import "fmt"

type person struct {
	name      string
	age       int
	sex, city string
	adult     bool
}

func main() {
	var wanger person
	fmt.Printf("wanger: %#v\n", wanger)
}
