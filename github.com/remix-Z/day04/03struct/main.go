package main

import "fmt"

//new和make都是给变量开辟内存空间
//new是给基本类型开辟内存的返回的是对应类型的指针
//make是给slice map chan开辟内存，返回的就是对应类型

//结构体 使用type和struct关键字
type person struct {
	name  string
	age   int
	hobby []string
	//score map[string]int
}

//简单使用
func test01() {
	//声明
	var p person
	//初始化
	p.name = "李明"
	p.age = 18
	p.hobby = []string{"篮球", "足球"}
	fmt.Println(p)

	//访问字段
	fmt.Println(p.name)
}

//匿名结构体
func test02() {

	//匿名结构体 多用于临时场景 基本只是用一次
	var s struct {
		x string
		y int
	}
	s.x = "hello world"
	s.y = 1
	fmt.Printf("%T %v\n", s, s)
}

func f1(x person) {
	x.age = 18 //只修改了副本的age，并不影响原结构体
}

func f2(x *person) {
	//(*x).age = 18 //根据内存地址修改原结构体
	x.age = 18 //语法糖 自动根据指针找原变量 go语言中不能对指针进行修改
}

//结构体是值类型的
func test03() {
	var p person
	p.name = "张三"
	p.age = 10
	fmt.Println(p)

	//值传递
	f1(p)
	fmt.Println(p)
	//指针传递
	f2(&p)
	fmt.Println(p)

	//创建指针类型结构体
	var p2 = new(person)
	fmt.Printf("%T %p\n", p2, p2) //*main.person
	p2.name = "李华"                //(*p2).name = "李华" 语法糖
	p2.age = 200

	//键值初始化
	//创建并初始化 加&就是指针 不加就是结构体
	var p3 = &person{
		name:  "王雷",
		age:   100,
		hobby: []string{"飞机", "大炮"},
	}
	fmt.Println(p3)

	//列表初始化
	//这种创建并初始化方式 结构体内部值的顺序和定义的顺序必须一致
	p4 := &person{
		"张飞",
		18,
		[]string{"火箭弹"},
	}
	fmt.Println(p4)
	fmt.Printf("%p,%p,%p\n", &(p4.name), &(p4.age), &(p4.hobby))

}

//结构体内存是连续的
type x struct {
	a int8 //8bit = 1byte
	b int8
	c int8
}

func test04() {
	//内存对齐
	x1 := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p,%p,%p\n", &(x1.a), &(x1.b), &(x1.c))
}

func main() {
	test01()

	test02()

	test03()

	test04()

}
