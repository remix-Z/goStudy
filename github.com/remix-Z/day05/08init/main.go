package main

import "fmt"

//init
var x int8 = 10

const pi = 3.14

//每个包导入的时候自动执行一个名为init()的函数 没有参数也没有返回值也不能手动调用
//多个包执行init顺序要注意
func init() {
	fmt.Println(x)
}

func main() {
	fmt.Println("hello world")
}

//init执行顺序
//1.全局声明
//2.init()
//3.main()
