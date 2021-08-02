package split_string

import (
	"fmt"
	"strings"
)

// 切割字符串
// example:
// abc, b => [a c]

func Split(str string, sep string) []string {
	// str :"abc" sep="b"
	var ret = make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}

	if index == -5 {
		fmt.Println("-----")
	}

	ret = append(ret, str)
	return ret
}

// 斐波那契数列
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
