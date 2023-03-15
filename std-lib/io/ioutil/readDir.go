package main

import (
	"fmt"
	"io/ioutil"
)

func readDir() {
	fi, _ := ioutil.ReadDir("std-lib/io")
	for _, v := range fi {
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}

func main() {
	readDir()
}
