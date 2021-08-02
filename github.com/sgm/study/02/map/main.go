package main

import (
	"fmt"
)

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        //还没有初始化到内存中
	m1 = make(map[string]int, 10) //要估算好map容量，避免在程序运行过程中在动态的扩容
	m1["理想"] = 18
	m1["上上"] = 35
	fmt.Println(m1)
	fmt.Println(m1["理想"])
	fmt.Println(m1["sad"]) // 如果不存在的key值 就取0值

	//v, ok := m1["ss"]
	//if !ok {
	//	fmt.Println("查无此人")
	//	os.Exit(1)
	//}
	//
	//fmt.Println(v)

	// map的遍历
	for k, v := range m1 {
		fmt.Println(k)
		fmt.Println(v)
	}

	// 删除一个不存在的则无任何影响
	delete(m1, "上上")
	fmt.Println(m1)

}
