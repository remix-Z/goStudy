package main

import "fmt"

//指针
func main() {
	//只有两个操作符
	//&:取地址，*:解引用
	a := 18
	b := [3]int{1, 2, 3}
	//fmt.Println(&a, &b)	//这样打不出来b的地址，打印出来的是&[1,2,3]
	fmt.Println(&a)
	fmt.Printf("%p", &b)

	// p := &a
	p := &b
	fmt.Printf("Type:%T\n", p) //*int类型 跟c一样
	fmt.Println(*p)

	//小问题来了
	// var p1 *int
	// p1 = 10
	// fmt.Println(*p1)

	var p1 *int
	fmt.Println(p1) //nil 只有make new出来的指针是有内存地址的
	p2 := new(int)
	fmt.Println(p2)
	*p2 = 100 //而不能*p1 = 100

	//make只能用于 slice map chan类型的创建,返回值是其本身而不是指针
	//new和make都是申请内存，new一般很少用 返回的是对应类型的指针，而make返回的是上述三种类型的引用

}
