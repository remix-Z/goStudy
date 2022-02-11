package main

import "fmt"

//recover必须搭配defer
//defer一定要在可能引发panic之前使用

//panic recover
func funcA() {
	fmt.Println("A")
}

func funcB() {
	//假设b出现了严重的错误
	//panic recover

	//场景应用 具体问题具体分析
	var a = 10
	//打开一个数据库连接 如果失败直接panic	假设0就是连接失败
	if a == 0 {
		defer func() {
			err := recover()
			fmt.Println(err)
			fmt.Println("释放数据库连接")
		}()
		panic("出现了严重的错误") //程序崩溃退出
	}

	fmt.Println("B")
}

func funcC() {
	fmt.Println("C")
}

func main() {
	funcA()
	funcB()
	funcC()
}
