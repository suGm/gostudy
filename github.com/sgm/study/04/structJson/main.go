package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "xxx",
		Age:  11,
	}

	b, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("Marshal fail")
		return
	}

	fmt.Printf("%#v\n", string(b)) // 序列化结构体至json

	// 反序列化
	var p2 person
	str := `{"name":"11","age":22}`
	json.Unmarshal([]byte(str), &p2) //传指针是为了能在json.Unmarshal内部修改p2的值
	fmt.Printf("%#v\n", p2)

}
