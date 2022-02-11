package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

//通过反射将配置文件的内容读到结构体变量中 类似JSON格式化

//配置文件结构体
type mysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type redisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Database string `ini:"database"`
	Password string `ini:"password"`
}

type Config struct {
	mysqlConfig `ini:"mysql"`
	redisConfig `ini:"redis"`
}

/*
	1. 参数校验
		1.1. 需要在函数中对参数赋值，所以参数必须是指针类型(结构体指针)
	2. 一行一行读取文件中的数据
	3. 如果是注释跳过
	4. 以'['开头表示节
	5. 如果不是'['开头则以'='分割
*/
func loadIni(fileName string, data interface{}) (err error) {
	t := reflect.TypeOf(data)
	//v := reflect.ValueOf(data)
	//校验传入参数是否为结构体指针类型 其实可有可无(自己调用的函数)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		err = errors.New("import parameter type must be *struct")
		return
	}

	//读文件
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//string(b) 将字节类型的文件内容转换成字符串
	lineSlice := strings.Split(string(b), "\r\n") //文件结尾\r\n

	//var structName string //用于记录section对应结构体
	//var fieldType reflect.StructField
	//分行读
	for _, line := range lineSlice {
		//去掉字符串首尾空格
		line = strings.TrimSpace(line)
		//如果是注释跳过(具体以注释开头为主# // ;)
		if strings.HasPrefix(line, "#") {
			continue
		}

		//如果以[开头 ]结尾 表示对应的section
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			sectionName := strings.TrimSpace(line[1 : len(line)-1]) //取出[]中间的section
			// //确定结构体类型
			// if slice == "mysql" {
			// 	var config1 mysqlConfig
			// 	data = config1
			// } else if slice == "redis" {
			// 	var config2 redisConfig
			// 	data = config2
			// }

			//根据字符串section去data里根据反射找对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				//fieldType = field
				if sectionName == field.Tag.Get("ini") {
					//structName = field.Name
					break //找到直接跳出循环吧	跳不跳都行吧
				}
			}

		} else {
			//根据structName去data里把对应的嵌套结构体取出
			//m := make(map[string]int, 128)
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//fmt.Println(line)
			fmt.Println(key, value)
			// sValue := v.Elem().FieldByName(structName)
			// sType := sValue.Type()
			//fmt.Println(sType)
			//fmt.Println(structObj.Type().Name())

			// var fieldName string
			// for i := 0; i < sValue.NumField(); i++ {
			// 	field := sType.Field(i)
			// 	//找到了对应的字段
			// 	if field.Tag.Get("ini") == key {
			// 		fieldName = field.Name
			// 		break
			// 	}
			// }

			//赋值
			//fileObj := sValue.FieldByName(fieldName)
			//fmt.Println(fieldName, fieldType.Type.Kind())
		}

	}

	//
	// fmt.Println(b[1])
	// fmt.Printf("%T", b[1])

	//因为传进来的是内存地址 所以要用到Elem()
	//{Address  string ini:"address" 0 [0] false} true
	//fmt.Println(t.Elem().FieldByName("Address"))
	//t.Elem().Field(0).Name = string(b[1])

	// for i := 0; i < t.Elem().NumField(); i++ {
	// 	//field := t.Elem().Field(i)
	// 	//fmt.Println(field.Name, field.Type, field.Index, field.Tag)

	// }

	// rd, err1 := ioutil.ReadFile("./test.ini")
	// if err1 != nil {
	// 	return
	// }
	// t1 := reflect.TypeOf(string(rd))
	// fmt.Println(t1)
	// //FieldByName等一系列方法 必须是struct类型才能调用,要不会panic
	// fmt.Println(t1.String())
	return
}

func main() {
	var cfg Config
	err := loadIni("./test.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed,err:%v", err)
		return

	}
	//fmt.Println(tmp)
}
