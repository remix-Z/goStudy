package main

import "fmt"

//切片
func main() {

	//定义
	var s1 []int
	var s2 []string
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1)) //刚开始就是空
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"北京", "上海", "深圳"}
	fmt.Println(s1, s2)

	//切片的长度和容量
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))

	//由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	a2 := a1[0:4] //左包含 右不包含 1, 3, 5, 7
	a3 := a1[1:6] // 3, 5, 7, 9, 11
	fmt.Println(a2)
	fmt.Println(a3)
	a4 := a1[:4] //1, 3, 5, 7
	a5 := a1[3:] //7, 9, 11, 13
	a6 := a1[:]  //1, 3, 5, 7, 9, 11, 13
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(a6)

	//切片指向了底层的数组，切片容量就是从指向的元素到数组最后的容量大小
	fmt.Printf("a4 len:%d,cap:%d\n", len(a4), cap(a4))
	fmt.Printf("a5 len:%d,cap:%d\n", len(a5), cap(a5))
	fmt.Printf("a6 len:%d,cap:%d\n", len(a6), cap(a6))

	//切片再切片
	a7 := a5[:3]
	fmt.Printf("a7 len:%d,cap:%d\n", len(a7), cap(a7))
	//更改底层数组，切片也会变化
	
	fmt.Println(a4)
	a1[0] = 100000
	fmt.Println(a4)
}
