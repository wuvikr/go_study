package main

import (
	"fmt"
	"log"
	"os"
)

// 创建文件
func createFile() {
	f, err := os.Create("./a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 读取文件
func readFile() {
	b, err := os.ReadFile("./std-lib/os/test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("b: %v\n", string(b[:])) // 文件传输以字节数组的方式
	}
}

// 写文件
func writeFile() {
	err := os.WriteFile("./std-lib/os/test.txt", []byte("hello world!"), os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 重命名文件
func renameFile() {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 创建目录
func createDir() {
	// err := os.Mkdir("a", os.ModePerm)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }

	// 创建多级目录
	err2 := os.MkdirAll("a/b/c", os.ModePerm)
	if err2 != nil && !os.IsExist(err2) {
		log.Fatal(err2)
	}
}

// 删除文件和目录
func remove() {
	// 删除文件
	// err := os.Remove("a.txt")
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }

	// 删除目录（递归删除）
	err2 := os.RemoveAll("a")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

// 获取当前工作目录路径 = pwd
func getWorkDir() {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)
}

func main() {
	// createFile()
	// createDir()
	// remove()
	// getWorkDir()
	// readFile()
	writeFile()
}
