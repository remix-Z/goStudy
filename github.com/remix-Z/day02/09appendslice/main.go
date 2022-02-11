package main

import "fmt"

//append() 为切片追加元素
func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1 V:%v,len:%d,cap:%d\n", s1, len(s1), cap(s1))
	//s1[3] = "广州"	//错误的写法 out of range
	s1 = append(s1, "广州")
	fmt.Printf("s1 V:%v,len:%d,cap:%d\n", s1, len(s1), cap(s1))

	ss := []string{"福建", "沈阳", "苏州"}
	s1 = append(s1, ss...) //ss是切片类型，第二个参数应该是字符串，'ss...'表示吧ss切片拆开
	fmt.Printf("s1 V:%v,len:%d,cap:%d\n", s1, len(s1), cap(s1))
}
