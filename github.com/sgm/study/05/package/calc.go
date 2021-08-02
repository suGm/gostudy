package calc

import "fmt"

// 包中的标识符（变量名、函数名、结构体、接口等）如果首字母是小写的只能在当前这个包使用
// 首字母大写的才能在其他包被调用

// 当被引用导入时自动调用 类似构造函数 没有参数也没有返回值
func init() {
	fmt.Println("自动调用")
}

func Add(x, y int) int {
	return x + y
}
