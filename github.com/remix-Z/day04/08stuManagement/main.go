package main

import (
	"fmt"
	"os"
)

/*
	学生管理系统
	可以查看/新增/删除/修改 学生

*/
var (
	allStudent map[int64]*Student
)

type Student struct {
	name  string
	numId int64
}

func showAllStudent() {
	//打印所有学生
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}

func createStudent(name string, numId int64) Student {
	return Student{
		name:  name,
		numId: numId,
	}
}

func addStudent() {
	//创建一个新学生
	var (
		id   int64
		name string
	)
	fmt.Printf("请输入学生编号:")
	fmt.Scanln(&id)
	fmt.Printf("请输入学生姓名:")
	fmt.Scanln(&name)
	newStu := createStudent(name, id)

	//追加到allStudent中
	allStudent[id] = &newStu
}

func deleteStudent() {
	var deleteId int64

	fmt.Println("请输入删除的编号:")
	fmt.Scanln(&deleteId)
	delete(allStudent, deleteId)
}

func main() {
	allStudent = make(map[int64]*Student, 50)
	// 1. 打印菜单
	// 2. 等待用户选择操作
	// 3. 执行对应的函数
	for {
		fmt.Println("***  欢迎光临  ***")
		fmt.Println("***学生管理系统***")
		fmt.Println(`***1.查看所有学生***
***2.新增学生  ***
***3.删除学生  ***
***4.退出	 ***`)
		fmt.Print("请输入您想要的操作:")

		var input int
		fmt.Scanln(&input)

		// if input == 1 {
		// 	fmt.Println("***1.查看所有学生***")
		// } else if input == 2 {
		// 	fmt.Println("***2.新增所有学生***")
		// } else if input == 3 {
		// 	fmt.Println("***3.删除所有学生***")
		// }
		switch input {
		case 1:
			//fmt.Println("***1.查看所有学生***")
			showAllStudent()
		case 2:
			//fmt.Println("***2.新增学生  ***")
			addStudent()
		case 3:
			//fmt.Println("***3.删除学生  ***")
			deleteStudent()
		case 4:
			//fmt.Println("***4.退出	 ***")
			os.Exit(1)
		default:
			fmt.Println("请重新输入")
			continue
		}
	}
}
