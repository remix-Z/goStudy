package main

import "fmt"

func main() {

	//获取用户输入
	var s string
	fmt.Print("请输入:")
	fmt.Scan(&s)
	fmt.Println(s)

	var name string
	var age int
	fmt.Println("请输入姓名和年龄")
	fmt.Scanf("%s %d\n", &name, &age) //只获取一次输入
	fmt.Println(name, age)

	var name02 string
	var age02 int
	fmt.Scanln(&name02, &age02)

}
