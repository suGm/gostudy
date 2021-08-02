package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	// time.Unix()
	ret := time.Unix(1623231283, 0)
	fmt.Println(ret)
	// 一秒
	fmt.Println(time.Second)
	// now + 1小时
	fmt.Println(now.Add(time.Hour))
	// 定时器
	//timer := time.Tick(time.Second)
	//for t := range timer {
	//	fmt.Println(t)
	//}

	// 格式化时间 把语言中的时间对象 转换成字符串类型的时间
	fmt.Println(now.Format("2006/01"))

	// 毫秒
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))

	// 按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2000-08-03")
	if err != nil {
		fmt.Println("parse time failed")
		return
	}

	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	nextYear, err := time.Parse("2006-01-02", "2020-08-03")

	if err != nil {
		fmt.Println("err3")
		return
	}
	//nextYear = nextYear.UTC()
	d := now.Sub(nextYear)
	fmt.Println(d)

	//time.Sleep(time.Second)
	//
	//fmt.Println("111")
	f2()
}

// 时区
func f2() {
	now := time.Now()
	fmt.Println(now)
	// 明天的这个时间
	// 按照指定格式去解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05", "2020-06-10 10:36:00")
	// 按照东八区的时区和格式去解析一个字符串格式的时间
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load loc failed")
		return
	}
	// 按照指定时区解释时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-06-10 10:36:00", loc)
	if err != nil {
		fmt.Println("parse time failed 2")
		return
	}

	fmt.Println(timeObj)
	// 时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}
