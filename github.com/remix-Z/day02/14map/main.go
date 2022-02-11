package main

import (
	"fmt"
)

//map
func main() {

	var m1 map[string]int

	//需要初始化开辟空间后才能使用
	m1 = make(map[string]int, 10) //估算好map的容量，避免动态扩容
	m1["李明"] = 18
	m1["张飒"] = 20
	m1["王五"] = 9

	fmt.Println(m1)

	value, ok := m1["王五"]
	if !ok {
		fmt.Printf("查无此人")
	} else {
		fmt.Printf("王五的value是:%d\n", value)
	}

	//map遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//只遍历 key
	for k1 := range m1 {
		fmt.Println(k1)
	}
	//只遍历 value
	for _, v1 := range m1 {
		fmt.Println(v1)
	}

	//删除delete
	delete(m1, "赵四")
	fmt.Println(m1)
	delete(m1, "王五")
	fmt.Println(m1)

}
