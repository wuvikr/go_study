package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

// student的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// 学生管理类型
type studentMgr struct {
	allStudents []*student
}

// studentMar的构造函数
func newstudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 1. 添加学生
func (s *studentMgr) addStudent(Stu *student) {
	s.allStudents = append(s.allStudents, Stu)
}

// 2. 编辑学生
func (s *studentMgr) modifyStudent(Stu *student) {
	for i, v := range s.allStudents {
		if Stu.id == v.id {
			s.allStudents[i] = Stu
			return
		}
	}

	fmt.Printf("输入的信息有误，系统中未找到学号为 %d 的学生", Stu.id)
}

// 3. 展示学生
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d， 姓名：%s， 班级：%s\n", v.id, v.name, v.class)
	}
}
