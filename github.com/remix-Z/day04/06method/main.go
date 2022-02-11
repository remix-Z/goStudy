package main

import (
	"fmt"
)

//标识符:变量名 函数名 雷兴明 方法名
//Go语言中如果标识符首字母大写 就表示其对外部包可用 (暴露的,公有的)
type Dog struct {
	name string
	age  int
}

//构造函数
func newDog(name string, age int) Dog {
	return Dog{
		name: name,
		age:  age,
	}
}

//方法 方法是作用域特定类型的函数
//接受者表示的是调用该方法的具体类型变量,多用于首字母小写
//(d Dog)这个叫接收者
func (d Dog) wang() {
	fmt.Printf("%s:汪汪汪", d.name)
}

//++ 需要修改值的话需要使用指针接收者
func (d *Dog) increase() {
	d.age++
}

func main() {
	d1 := newDog("小黑", 2)
	d1.wang()
	fmt.Println(d1)
	d1.increase()
	fmt.Println(d1)
}

//
