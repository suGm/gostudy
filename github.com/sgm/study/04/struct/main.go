package main

import "fmt"

// 自定义类型和类型别名

type myInt int     // 自定义类型
type yourInt = int // 类型别名

// 结构体 是一个连续的内存
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

type x struct {
	a int8
	b int8
	c int8
}

// 匿名字段
type person3 struct {
	string
	int
}

// 结构体是值类型
type person1 struct {
	name, gender string
}

func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m yourInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	// 申明一个person类型的变量p
	var p person
	// 通过字段赋值
	p.name = "周"
	p.age = 18
	p.gender = "男"
	p.hobby = []string{"篮球", "足球"}

	fmt.Printf("%T\n", p)

	var p2 person
	p2.name = "理想"
	p2.age = 18
	fmt.Printf("type:%T value:%v\n", p2, p2)

	// 匿名结构体 多用于一些临时场景
	var s struct {
		name string
		age  int
	}

	s.name = "heihei"
	s.age = 100
	fmt.Printf("type:%T value:%v\n", s, s)

	// 结构体是值类型
	var p1 person1
	p1.name = "p1"
	p1.gender = "男"

	f(&p1)
	fmt.Printf("%v\n", p1)

	var p3 = new(person)
	fmt.Printf("%T\n", p3)
	fmt.Printf("%p\n", p3)

	var a int
	a = 100
	b := &a
	fmt.Printf("type a:%T type b:%T\n", a, b)
	// 将a的十六进制内容地址打印出来
	fmt.Printf("%p\n", &a) // a的内存地址
	fmt.Printf("%p\n", b)  // b的值
	fmt.Printf("%p\n", &b) // b的内存地址

	var p4 = person{
		name:   "aaa",
		gender: "nv",
	}

	fmt.Printf("%p\n", p4)

	p5 := person1{
		"bbb",
		"gender",
	}
	fmt.Printf("%T\n", p5)

	p6 := person1{
		name: "bbb",
	}
	fmt.Printf("%T\n", p6)

	m1 := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}

	fmt.Printf("%p\n", &(m1.a))
	fmt.Printf("%p\n", &(m1.b))
	fmt.Printf("%p\n", &(m1.c))
}

// go 语言传参默认永远是传值类型 永远是拷贝 除非传指针
func f(x *person1) {
	x.gender = "女"
}
