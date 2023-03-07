package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取当前进程PID
	fmt.Printf("os.Getpid(): %v\n", os.Getpid())

	// 获取父进程PID
	fmt.Printf("os.Getppid(): %v\n", os.Getppid())

	// 设置新进程属性
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Env:   os.Environ(),
	}

	// 开始一个新进程
	p, err := os.StartProcess("C:\\Users\\evwu\\AppData\\Local\\Programs\\Microsoft VS Code\\Code.exe", []string{"C:\\Users\\evwu\\AppData\\Local\\Programs\\Microsoft VS Code\\Code.exe", "C:\\Users\\evwu\\Desktop\\a.txt"}, attr)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("p: %v\n", p)

	// 通过进程PID寻找进程
	p2, _ := os.FindProcess(p.Pid)
	fmt.Printf("p2: %v\n", p2)
}
