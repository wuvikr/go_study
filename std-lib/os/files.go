package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// 打开文件，获取文件操作对象
func openFile() {
	// 只读方式打开文件
	// f, err := os.Open("./std-lib/os/test.txt")

	// 打开文件并指定权限
	f, err := os.OpenFile("./std-lib/os/test.txt", os.O_RDWR|os.O_CREATE, 755)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	}
}

// 创建文件
func createFile() {
	// 等价于 OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	// f, _ := os.Create("./std-lib/os/test2.txt")
	// fmt.Printf("f.Name(): %v\n", f.Name())

	// 第一个参数为临时文件路径，如果为空，则使用系统默认临时文件目录
	// 第二个参数为文件名，真正的文件名后面会加上一个随机字符串
	// 如果第二个参数包含"*"，则随机字符串替换最后的"*"，一般用于添加文件名后缀
	//f2, err := os.CreateTemp("", "example")
	f2, err := os.CreateTemp("", "temp.*.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("f2.Name(): %v\n", f2.Name())
	// 输出：	f2.Name(): C:\Users\evwu\AppData\Local\Temp\example3092511300
	//			f2.Name(): C:\Users\evwu\AppData\Local\Temp\temp.2826460038.txt
}

// 读取文件内容
func readFile() {
	f, _ := os.Open("./std-lib/os/test.txt")
	// 循环读取数据
	for {
		b := make([]byte, 5) // 构建缓冲字节数组
		// f.Seek(3, 0) // 从第三个字节开始读取数据
		n, err := f.Read(b) //读取数据，返回值为读取到的字节数
		fmt.Printf("string(b): %v\n", string(b))
		fmt.Printf("n: %v\n", n)
		if err != nil && err == io.EOF { // 中断循环
			fmt.Printf("err: %v\n", err)
			break
		}
	}
	f.Close()
}

// 读取目录下其他目录
func readDir() {
	f, _ := os.Open("./std-lib")
	de, _ := f.ReadDir(-1)
	for _, v := range de {
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
	f.Close()
}

// 写字节数组
func write() {
	f, _ := os.OpenFile("./std-lib/os/test.txt", os.O_RDWR|os.O_APPEND, 0775)
	f.Write([]byte("welcome to golang."))
	f.Close()
}

// 写字符串
func writeString() {
	f, _ := os.OpenFile("./std-lib/os/test.txt", os.O_RDWR|os.O_TRUNC, 0775)
	f.WriteString("\nis this our home?")
	f.Close()
}

func main() {
	// openClose()
	// createFile()
	// readFile()
	// readDir()
	// write()
	writeString()
}
