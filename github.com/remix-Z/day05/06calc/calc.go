package calc

//包中的标识符(变量名/函数名/结构体名/接口等)如果首字母是小写的，表示私有(只能在当前的包中使用)
//如果想在其他包中使用 首字母大写 外部可见(调用)
//func add(x,y)int{return x+y}
func Add(x, y int) int {
	return x + y
}
