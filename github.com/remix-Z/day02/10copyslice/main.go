package main

import (
	"fmt"
	"sort"
)

func main() {
	a1 := []int{1, 2, 3, 5, 7, 8, 10}
	a2 := a1
	var a3 = make([]int, 7)
	copy(a3, a1)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	//go语言中没有删除切片的操作，只能靠切片的特性来实现
	//例如要删除a1中的第三位元素
	a1 = append(a1[:2], a1[3:]...)
	fmt.Println(a1, a2, a3)

	fmt.Println()
	//总结一下
	x1 := [...]int{1, 2, 3, 4, 5}
	s1 := x1[:] //切片
	s1[0] = 100
	fmt.Println(x1, s1) //对切片操作，会更改底层数组

	s1 = append(s1[:2], s1[3:]...) //也会对底层数组进行操作
	//s1 = append(s1[:2], s1[2:len(s1)-1]...) //下面一行的操作 这行是错的
	//s1 = x1[:len(s1)-1] 或者 s1 = x1[1:]
	//想要去掉首尾那直接在声明那块操作就可以了 何必用到append呢，陷入自己的思维误区了 s1 := x1[1:]不就好了，用什么append
	fmt.Println(x1, s1)

	s2 := []int{200, 5, 4, 3}
	s1 = append(s2[:2], s2[3:]...) //s1指向的底层数组变了
	fmt.Println(x1, s1)

	//练习 坑
	// var a = make([]string, 5, 10)
	// for i := 0; i < 10; i++ {
	// 	a = append(a, fmt.Sprintf("%v", i))
	// }
	// fmt.Println(a)

	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	//引申 sort
	array := [...]int{9, 7, 3, 8, 6, 2, 1}
	slice := array[:]
	sort.Ints(slice) //也会对底层数组array进行操作
	fmt.Println(slice, array)
}
