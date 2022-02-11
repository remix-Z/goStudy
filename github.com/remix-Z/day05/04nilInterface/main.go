package main

import "fmt"

//空接口 interface {} 没有必要起名字

//当你不知道别人传进来什么类型的时候
//所有的类型都实现了空接口 任意类型都能传进来
func main() {
	//不知道什么类型的时候 用空接口interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)

	m1["name"] = "鲲鲲"
	m1["age"] = 18
	m1["married"] = false
	m1["hobby"] = [...]string{"唱", "跳", "rap", "篮球"}

	fmt.Println(m1)

	showTime(m1)
	showTime(nil)
	showTime(true)
	showTime(1)
}

//空接口作为参数
func showTime(a interface{}) {
	//fmt.Println("show Time")
	fmt.Printf("type:%T value:%v\n", a, a)
}

//interface{}
//动态类型
//动态值

//类型断言 x.T
