package main

import (
	"encoding/json"
	"fmt"
)

//结构体与JSON 前后端分离
/*
{
	{
		"name" : "小明",
		"age" : 100
	}
}
*/

// 1.序列化 把Go语言中的结构体变量 ---> JSON格式的字符串
// 2.反序列化 JSON格式的字符串 ---> Go语言能够识别的结构体变量

/*
	Go语言通过首字母的大小写来控制访问权限。
	无论是方法，变量，常量或是自定义的变量类型，如果首字母大写，则可以被外部包访问，反之则不可以。
	如果希望，json化之后的属性名是小写字母的，可以使用struct tag。如下 `json:"xxxx"`
*/
type person struct {
	Name string //首字母大写
	Age  int    //首字母不大写 最后拿不到因为json.Marshal调了json的包而不是你main的包
	//age  int
	Age2 int `json:"age2"` //如果前段要求首字母必须小写就要用这种方式 `json:"xxxx"`
}

func main() {
	p1 := person{
		Name: "小明",
		Age:  100,
		Age2: 18,
	}

	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed ,err:%v\n", err)
		return
	}

	fmt.Println(string(b))

	//反序列化
	str := `{"name":"李华","age":25,"age2":20}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传地址是为了能在json.Unmarshal内部修改p2的值
	fmt.Println(p2)
}
