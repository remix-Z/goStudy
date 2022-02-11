package main

import "fmt"

//构造函数

type person struct {
	name string
	age  int
}

//构造函数
//返回的是结构体或是结构体指针 需要判断资源(内存)消耗
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//当结构体过大时 需要返回结构体指针
func newPerson02(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("小明", 16)
	p2 := newPerson("丽丽", 20)
	p3 := newPerson02("王思", 21)

	fmt.Println(p1, p2, p3)
}
