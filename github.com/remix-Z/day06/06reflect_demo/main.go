package main

import (
	"fmt"
	"reflect"
)

//反射 reflect包
/*
	Golang提供了一种机制，在编译时不知道类型的情况下，可更新变量、运行时查看值、调用方法以及直接对他们的布局进行操作的机制，称为反射。
	反射主要:
		1.获取变量内部信息
			类型				作用
			reflect.ValueOf()	获取输入参数接口中的数据的值，如果为空则返回0 <- 注意是0
			reflect.TypeOf()	动态获取输入参数接口中的值的类型，如果为空则返回nil <- 注意是nil
		2.struct的反射
		3.通过反射设置变量的值
		4.IsNil() IsValid()


*/
//获取变量内部类型
func reflectTypeOf(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("TYPE:%v\n", v)
}

//结构体 类型
func reflectKind(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("TYPE.NAME:%v TYPE.KIND:%v\n", v.Name(), v.Kind())
}

func reflectValueOf(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind() //值的类型种类
	fmt.Printf("VALUE:%v KIND:%v\n", v, k)
}

//通过反射设置变量的值 值拷贝
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本 reflect包会引发panic
	}
}

//指针
func reflectSetValueNew(x interface{}) {
	v := reflect.ValueOf(x)
	//反射中通过Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200) //修改的是副本 reflect包会引发panic
	}
}

type Cat struct{}

func main() {
	var a float64 = 1.6
	reflectTypeOf(a)
	reflectValueOf(a)

	var b int64 = 100
	reflectTypeOf(b)

	var c *int32
	reflectTypeOf(c)

	var d = Cat{}
	reflectTypeOf(d) //main.Cat
	reflectKind(d)   //Cat和struct

	//设置值
	//reflectSetValue(&b)
	reflectSetValueNew(&b)
	fmt.Println(b)

	test01()
	testStruct()

}

func test01() {
	var p *int
	//IsNil()
	fmt.Println("var a *int IsNil:", reflect.ValueOf(p).IsNil())
	fmt.Println("var a *int IsNil:", reflect.ValueOf(p).IsValid())

	//实例化一个匿名结构体
	b := struct{}{}
	//尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	//尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())

	//map
	//var c map[string](int)
	c := map[string](int){}
	//尝试从map中查找"abc"的键
	fmt.Println("map中不存在的键:", reflect.ValueOf(c).MapIndex(reflect.ValueOf("abc")).IsValid())

}

//结构体反射
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func testStruct() {
	stu1 := student{
		Name:  "王五",
		Score: 100,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())

	//通过for循环遍历结构体的所有字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:[%s] index:[%d] type:[%v] json tag:[%v]\n", field.Name, field.Index, field.Type, field.Tag)
	}

	//通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag)
	}
}
