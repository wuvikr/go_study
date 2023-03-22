package main

import (
	"fmt"
	"os"
)

// 学生管理系统

func showMenu() {
	fmt.Println("欢迎来到学生信息管理系统")
	fmt.Println("1. 添加学生")
	fmt.Println("2. 编辑学生信息")
	fmt.Println("3. 展示所有学生信息")
	fmt.Println("4. 退出系统")
}

func getInput() *student {
	var (
		id    int
		name  string
		class string
	)

	fmt.Println("请按要求输入学生信息！")
	fmt.Print("请输入学生的学号: ")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学生的姓名: ")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学生的班级: ")
	fmt.Scanf("%s\n", &class)

	stu := newStudent(id, name, class)
	return stu
}

func main() {

	sm := newstudentMgr()

	for {
		// 1. 打印系统菜单
		showMenu()

		// 2. 等待用户选择要执行的选项
		var input int
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户输入的是：", input)

		// 3. 执行用户选择的动作
		switch input {
		case 1:
			// 添加学生
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			// 编辑学生
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			// 展示所有学生
			sm.showStudent()
		case 4:
			os.Exit(4)
		}
	}
}
