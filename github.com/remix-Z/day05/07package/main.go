package main

//包的导入
import (
	"fmt"
	//会从GOROOT和$(GOPATH)/src/下找 前面的calc是别名(文件夹名和文件名不一致时)
	calc "github.com/remix-Z/day05/06calc"
)

func init() {
	fmt.Println("自动执行!")
}

func main() {
	//外部调用首字母需大写，要不找不到
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}
