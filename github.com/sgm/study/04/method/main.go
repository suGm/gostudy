package main

import "fmt"

// 方法

// go语言中如果标识符首字母是大写的，就表示对外是可见的(能够被别的包调用的)
// 不能给别的包里面的类型定义方法，只能给自己的包里定义

// Dog 这是一个狗的结构体
type Dog struct {
	name string
	age  int
}

// 构造函数
func newDog(name string) Dog {
	return Dog{
		name: name,
	}
}

// 方法区别于函数 方法是作用于特定类型的函数
// (d dog)被称为接受者，接受者表示的是调用该方法的具体类型变量。多用类型名首字母小写表示
func (d *Dog) wang() {
	fmt.Printf("%s:汪汪汪\n", d.name)
}

func (d *Dog) guonian() {
	d.age += 1
}

func main() {
	var d = newDog("a")
	d.wang()
	d.guonian()
	fmt.Println(d.age)
}
