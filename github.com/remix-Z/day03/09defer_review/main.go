package main

import "fmt"

//defer面试题
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret

}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

//index a b ret
//defer calc里面的calc会正常执行，外层defer延迟执行,但内部传参已经记录
//10,1,2,3	//calc("10",1,2)
//20,0,2,2	//calc("20",0,2)
//2,0,2,2	//defer calc("2",0,2)	//此时a = 0已经被修改
//1,1,3,4	//defer calc("1",1,3)	//此时再第16行调用时已经记录a是1，而不是下面的a
//defer就近取值(向上)

//执行顺序
//1. a := 1
//2. b := 2
//3. defer calc("1",1,calc("10",1,2))
//4. calc("10",1,2)	//打印 10 1 2 3
//5. a = 0
//6. defer calc("2",0,calc("20",0,2))
//7. calc("20",0,2)	//打印 20 0 2 2
//8. b = 1
//9. calc("2",0,2)	//打印 2 0 2 2
//10. calc("1",1,3) //打印 1 1 3 4
