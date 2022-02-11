package main

import "fmt"

//map和slice的组合 ***初始化***
func main() {
	//元素类型为map的切片
	s1 := make([]map[int]string, 10, 10)

	//s1第一个map切片的，赋予key为100的value为"A"
	//s1[0][100] = "A"	map没有初始化 不能直接用
	//fmt.Println(s1)

	s1[0] = make(map[int]string, 1)
	s1[0][10] = "A"
	fmt.Println(s1)

	//值为切片的map
	var s2 = make(map[string][]int, 10)
	s2["上海"] = []int{10, 20, 30}
	fmt.Println(s2)

}
