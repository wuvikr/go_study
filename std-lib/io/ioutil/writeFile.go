package main

import "io/ioutil"

func writeFile() {
	ioutil.WriteFile("./a.txt", []byte("hello golang!"), 0644)

}

func main() {
	writeFile()
}
