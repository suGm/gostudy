package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

type float641 float64

func reflecType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v)
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200) // 修改的是副本, reflect会引发panic
	}
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0、参数的校验
	// 0.1、传进来的data参数必须是指针类型(因为需要在函数中对其赋值)
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer") // 创建一个错误
		return err
	}
	// 0.2、传进来的data参数必须是结构体类型指针（因为配置文件中各种键值对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a struct")
		return err
	}
	// 1、读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("err 1")
		return err
	}

	lineSlice := strings.Split(string(b), "\n") // 将文件内容转换成字符串
	fmt.Printf("%#v\n", lineSlice)
	var structName string
	// 2、一行一行得读数据
	for idx, line := range lineSlice {
		// 去除字符串首尾空格
		line = strings.TrimSpace(line)
		// 2.1 如果是注释就跳过
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "[") {
			// 2.2 如果是[]开头的就表示是节(section)
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			// 把这一行首尾的[]去掉，取到中间的内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			// 根据字段字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应的结构体，把字段名记下来
					structName = field.Name
					fmt.Println("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 2.3 如果不是[开头就是=分割的键值对
			// 1、以等号分割这一行，等号左边是key，等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			//value := strings.TrimSpace(line[index+1:])
			// 2、根据structName去data里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的值信息
			sType := sValue.Type()                     // 拿到嵌套结构体的类型信息

			if sType.Kind() != reflect.Struct {
				fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			// 3、遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sType.NumField(); i++ {
				field := sType.Field(i) // tag信息时存储在类型信息中的
				if field.Tag.Get("ini") == key {
					// 找到对应的字段
					fieldName = field.Name
				}
			}
			// 4、如果key = tag，给这个字段赋值
			// 4.1 根据filedName去取出这个字段
			fileObj := sValue.FieldByName(fieldName)
			// 4.2 对其赋值
			fmt.Println(fieldName, fileObj.Type().Kind())
		}
	}

	return nil
}

func main() {
	//str := `{"name":"搜索", "age":10}`
	//var p person
	//json.Unmarshal([]byte(str), &p)
	//fmt.Println(p)
	//var a float64 = 3.14
	//reflecType(a)
	//var b float641 = 3.14
	//reflecType(b)
	//var c int64 = 100
	//reflectSetValue1(&c)
	//fmt.Println(c)
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cfg)
}
