package main

import (
	"flag"
	"fmt"
	"time"
)

// flag 获取命令行参数
func main() {
	//fmt.Printf("%#v\n", os.Args)
	//fmt.Println(os.Args[0], os.Args[1], os.Args[2])
	//fmt.Printf("%T\n", os.Args)

	// 创建一个标志位参数
	name := flag.String("name", "sgm", "请输入名称")
	age := flag.Int("age", 1111, "请输入真实年龄")
	married := flag.Bool("married", false, "婚否")
	cTime := flag.Duration("ct", time.Second, "结婚多久了")
	// 使用flag
	flag.Parse()
	fmt.Printf("%#v\n", *name)
	fmt.Printf("%#v\n", *age)
	fmt.Printf("%#v\n", *married)
	fmt.Printf("%#v\n", *cTime)

}
