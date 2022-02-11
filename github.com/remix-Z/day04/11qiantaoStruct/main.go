package main

import (
	"fmt"
)

//嵌套结构体
type Person struct {
	name Name
	age  int
}

//name包含姓 和 名
type Name struct {
	xing string
	ming string
}

func test01() {
	p1 := Person{
		name: Name{
			xing: "张",
			ming: "三",
		},
		age: 18,
	}

	fmt.Println(p1)
	fmt.Println(p1.name.xing)
	fmt.Println(p1.name.ming)
	fmt.Println(p1.age)
}

//如果使用匿名嵌套结构体，就有语法糖
type company struct {
	name    string
	address //匿名嵌套
}

type address struct {
	province string
	city     string
}

func test02() {
	c1 := company{
		name: "神奇公司",
		address: address{
			"辽宁",
			"沈阳",
		},
	}

	//语法糖
	fmt.Println(c1.province) //现在机子的结构体中查找字段 再从匿名的嵌套结构体中查找
	fmt.Println(c1.city)

	//但如果有多个匿名嵌套结构体字段冲突 就不能使用语法糖
	
}

func main() {
	test01()

	test02()

}
