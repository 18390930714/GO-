package main

import "fmt"

type Student struct {
	id    int //学号是唯一的
	name  string
	class string
}

// Student 构造函数
func NewStudent(id int, name, class string) *Student {
	return &Student{
		id:    id,
		name:  name,
		class: class,
	}
}

// 学员管理类型
type StudentMgr struct {
	allStudents []*Student
}

func NewStudentMgr() *StudentMgr {
	return &StudentMgr{
		allStudents: make([]*Student, 0, 100),
	}
}

// 添加学员信息
func (s *StudentMgr) AddStudent(newStu *Student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 修改学员信息
func (s *StudentMgr) ModifyStudent(newStu *Student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id {
			s.allStudents[i] = newStu
			return
		}
	}
	fmt.Printf("输入学员信息有误，系统中没有学号是：%d的学生\n", newStu.id)
}

// 展示学员信息
func (s *StudentMgr) ShowStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n", v.id, v.name, v.class)
	}
}
