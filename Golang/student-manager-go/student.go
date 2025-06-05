package main

import "fmt"

// Student 定义一个结构体描述学生信息
type student struct {
	name  string
	id    int
	class string
}

// 定义student类型的构造函数
func newStudent(name string, id int, class string) *student {
	return &student{
		name:  name,
		id:    id,
		class: class,
	}
}

// 学员管理的类型
type studentMgr struct {
	allStudents []*student //student类型的切片
}

// 这是studentMgr的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100), //初始化切片
	}
}

// 1.添加学员
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 2.编辑学员
func (s *studentMgr) editStudent(student *student) {
	for i, v := range s.allStudents {
		if v.id == student.id {
			s.allStudents[i] = student
			return
		}
	}
	//如果走到这里说明输入的学生没有找到
	fmt.Println("输入的学员信息有误,没有找到学员编号为%d的学员", student.id)
}

// 学员列表
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Println(v)
	}
}
