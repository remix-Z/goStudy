package main

import (
	"fmt"
	"unsafe"
)

// 内存对齐
// 二个原则：
// 1.具体类型，对齐值=min(编译器默认对齐值，类型大小Sizeof长度)
// 2.struct每个字段内部对齐，对齐值=min(默认对齐值，字段最大类型长度)
func test01() {

	//相同的成员，不同的排列顺序
	fmt.Println(unsafe.Sizeof(struct {
		i8  int8
		i16 int16
		i32 int32
	}{})) //8
	//最大32 就是4字节
	//那么排列就是 x _ x x | x x x x  ---> 8

	fmt.Println(unsafe.Sizeof(struct {
		i8  int8
		i32 int32
		i16 int16
	}{})) //32
	//排列就为 x _ _ _ | x x x x | x x _ _   ---> 12
}

type Test struct {
	i8  int8
	i16 int16
	i32 int32
}

// 功能			函数
// 获取对齐值	unsafe.Alignof(t.i16)
// 获取偏移值	unsafe.Offsetof(t.i16)

func test02() {
	var test = new(Test)

	//从0开始
	var i8 = (*int8)(unsafe.Pointer(test))
	*i8 = int8(10)

	//偏移int8 +1字节数
	var i16 = (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(test)) + uintptr(1) + uintptr(unsafe.Sizeof(int8(0)))))
	*i16 = int16(10)

	//偏移int8 +1 + int16字节数
	var i32 = (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(test)) + uintptr(1) + uintptr(unsafe.Sizeof(int8(0))) + uintptr(unsafe.Sizeof(int16(0)))))
	*i32 = int32(10)

	fmt.Println(*test)
	fmt.Println(unsafe.Sizeof(test))

	//有区别吗???	不是很懂
	test2 := Test{
		i8:  int8(10),
		i16: int16(10),
		i32: int32(10),
	}
	fmt.Println(test2)
	fmt.Println(unsafe.Sizeof(test2))
}

func main() {
	test01()

	test02()
}
