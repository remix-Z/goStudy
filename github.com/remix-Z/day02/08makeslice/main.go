package main

import (
	"fmt"
)

//通过make创建切片
func main() {
	//参数类型，长度，容量
	s1 := make([]int, 5, 10) //创建有一个长度为5，容量为10的int切片
	fmt.Printf("s1 V:%v,len:%d,cap:%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s1 V:%v,len:%d,cap:%d\n", s2, len(s2), cap(s2))

	//nil的切片是没有底层数组的

	//切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3
	fmt.Println(s3, s4)
	s3[0] = 100 //有点类似c里的指针，这个操作直接吧底层数组的值给改了
	fmt.Println(s3, s4)

	//切片遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}

	for _, v := range s4 {
		fmt.Println(v)
	}
}
