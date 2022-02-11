package main

import "fmt"

//给任意类型添加方法
//不能给别的包的类型添加方法 只能给自己的添加方法
//如果想添加 需要基于别人包里的类型 再创建自己的类型

type myInt int

func (m myInt) hello() {
	fmt.Printf("这是我的int:%d\n", m)
}

func main() {
	m := myInt(10)
	m.hello()
}
