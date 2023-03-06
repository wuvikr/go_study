package main

import "fmt"

type frog struct {
	num   int
	color string
	size  string
}

type fisher interface {
	swim()
}

type reptile interface {
	crawl()
}

func (frog) swim() {
	fmt.Println("frog can swim.")
}

func (frog) crawl() {
	fmt.Println("frog can crawl.")
}

func main() {
	var a = frog{
		1,
		"green",
		"small",
	}
	var i fisher
	i = a
	i.swim()
}
