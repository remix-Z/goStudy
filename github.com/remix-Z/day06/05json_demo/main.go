package main

import (
	"encoding/json"
	"fmt"
)
//反射 
//struct field tag `json:name` not compatible(不兼容) with reflect.StructTag.Get: bad syntax(语法) for struct tag valuestructtag
type person struct {
	Name string `json:"name"` //`json:name`	name得有双引号
	Age  int    `json:"age"`  //`json:age`
}

func main() {
	//json
	str := `{"name":"李旺","age":9000}`
	var p person
	json.Unmarshal([]byte(str), &p) //反射
	fmt.Println(p.Name, p.Age)
}
