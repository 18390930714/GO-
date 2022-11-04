package main

import (
	"fmt"
	"os"
)

// 打印系统菜单
// 1.添加学院信息
// 2.编辑学员信息
// 3.展示学院信息

func showMenu() {
	fmt.Println("欢迎进入学员系统")
	fmt.Println("1.添加学员信息")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出系统")
}

//获取用户键入的学员信息

func getInput() *Student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请输入学员信息：")
	fmt.Println("请输入学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Println("请输入名字：")
	fmt.Scanf("%s\n", &name)
	fmt.Println("输入班级：")
	fmt.Scanf("%s\n", &class)

	stu := NewStudent(id, name, class) //newStudent()函数返回的是指针
	return stu
}

func main() {
	sm := NewStudentMgr()
	for {
		showMenu()

		var input int
		fmt.Printf("请输入选项：")
		fmt.Scanf("%d\n", &input)
		switch input {
		case 1:
			//添加学员
			stu := getInput()
			sm.AddStudent(stu)
		case 2:
			//编辑学员
			stu := getInput()
			sm.ModifyStudent(stu)
		case 3:
			//展示所有学员
			sm.ShowStudent()
		case 4:
			//退出
			os.Exit(0)
		}

	}
}
