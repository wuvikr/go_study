package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content := []byte("This is a temporary file")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tmpfile.Name(): %v\n", tmpfile.Name())

	//defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
