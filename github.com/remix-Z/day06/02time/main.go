package main

import (
	"fmt"
	"time"
)

//时间
func main() {
	//time.Now()
	fmt.Println(time.Now())
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Date())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())

	//时间戳
	timeStampDemo()

	//time.Unix() 时间是从1970年开始
	ret := time.Unix(730771200, 0)
	fmt.Println(ret)

	//时间间隔
	fmt.Println(time.Second) //1s钟

	//Add now()+1h 当前时间加1小时
	fmt.Println(time.Now().Add(time.Hour))
	//fmt.Println(time.Now().Sub(time.Now()))

	//定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	//时间格式化
	formatDemo()

	//将按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "1993-02-27")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	fmt.Println("------------------")
	timeEdge()

}

func timeStampDemo() {
	now := time.Now()
	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano() //纳秒

	fmt.Println(timeStamp1)
	fmt.Println(timeStamp2)
}

//时间格式化
func formatDemo() {
	//Go语言中格式化时间模拟的不是常见的'Y-M-D H:M:S' 而是使用Go诞生的时间2006年1月2号15点04分(记忆口诀 2006 1 2 3 4)
	now := time.Now()
	//格式化时间模板 2006年1月2日15点04分 Mon Jan
	//24小时
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	//fmt.Println(now.Format("15:06 2001/01/02"))

	//12小时
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
}

//时区
func timeEdge() {
	now := time.Now() //获取本地时间
	fmt.Println(now)

	//明天的这个时间
	time.Parse("2006-01-02 15:04:05", "2022-01-18 15:19:00")

	//按照东八区的时间和格式解析一个字符串格式的时间
	//根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	//按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-01-18 15:19:00", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)

	//sub	时间对象相减
	sub := timeObj.Sub(now)
	fmt.Println(sub)
}
