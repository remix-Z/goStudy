package main

import (
	"fmt"
	"os"
)

type Student struct {
	name  string
	score int
}

//构造函数 创建一个学生
func createStudent(name string, score int) *Student {
	return &Student{
		name:  name,
		score: score,
	}
}

//一个学生管理系统
type StudentSystem struct {
	allStudent map[int]*Student
}

//新增学生
func (ss *StudentSystem) addStudent() {
	var id int
	fmt.Print("请输入想要新增的学生ID:")
	fmt.Scanln(&id)
	var name string
	fmt.Print("请输入想要新增的学生姓名:")
	fmt.Scanln(&name)
	newStudent := createStudent(name, 0)

	//追加到allStudent中
	ss.allStudent[id] = newStudent
}

//查看所有学生
func (ss *StudentSystem) showStudent() {
	for k, v := range ss.allStudent {
		fmt.Printf("id:%d 姓名:%s 分数:%d\n", k, v.name, v.score)
	}
}

//删除学生
func (ss *StudentSystem) deleteStudent() {
	fmt.Print("请输入想要删除的学生id:")
	var deleteId int
	fmt.Scanln(&deleteId)
	delete(ss.allStudent, deleteId)
}

func main() {
	//创建一个学生系统 容量为48
	var s1 StudentSystem
	s1.allStudent = make(map[int]*Student, 48)

	fmt.Println("新一代学生管理系统2.0")
	for {
		fmt.Println("******************")
		fmt.Println(`1.查看所有学生
2.新增学生
3.删除学生
4.退出`)
		fmt.Print("请输入您想要的操作:")
		var input int

		fmt.Scanln(&input)
		switch input {
		case 1:
			s1.showStudent()
		case 2:
			s1.addStudent()
		case 3:
			s1.deleteStudent()
		case 4:
			os.Exit(1)
		}
	}

}
