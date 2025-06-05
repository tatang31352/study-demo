package main

import (
	"fmt"
	"os"
)

//学员信息管理系统

//1.添加学员信息
//2.编辑学员信息
//3.展示所有学员信息

func showMenu() {
	fmt.Println("欢迎来到学院管理系统")
	fmt.Println("1.添加学员信息")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示学生列表")
	fmt.Println("4.退出学员系统")
}

// 获取用户输入的学员信息
func getInputStu() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请输入学员信息")
	fmt.Println("请输入学员学号：")
	fmt.Scanf("%d", &id)
	fmt.Println("请输入学员名字：")
	fmt.Scanf("%s", &name)
	fmt.Println("请输入学员班级：")
	fmt.Scanf("%s", &class)
	stu := newStudent(name, id, class)
	return stu
}

func main() {
	sm := newStudentMgr()
	for {
		//1.打印一个系统菜单
		showMenu()
		//2.等待用户选择要执行的选项
		var input int
		fmt.Println("请输入要操作的编号")
		fmt.Scanf("%d", &input)
		fmt.Println("您输入的编号是:", input)

		//3.执行相应动作
		switch input {
		case 1:
			stu := getInputStu()
			sm.addStudent(stu)
		case 2:
			stu := getInputStu()
			sm.editStudent(stu)
		case 3:
			sm.showStudent()
		case 4:
			os.Exit(0)
		}
	}
}
