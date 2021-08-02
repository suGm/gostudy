package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	fmt.Printf("left:%d\n", dispatchCoin())
}

func dispatchCoin() (left int) {
	e := []rune{'e', 'E'}
	i := []rune{'i', 'I'}
	o := []rune{'o', 'O'}
	u := []rune{'u', 'U'}

	sum := 0

	funcE := handleCoin(e, 1)
	funcI := handleCoin(i, 2)
	funcO := handleCoin(o, 3)
	funcU := handleCoin(u, 4)

	var sumUser int
	for _, user := range users {
		sumUser = 0
		sumUser += funcE(user)
		sumUser += funcI(user)
		sumUser += funcO(user)
		sumUser += funcU(user)
		fmt.Printf("%s获取金币数:%d\n", user, sumUser)
		sum += sumUser
	}

	return coins - sum
}

func handleCoin(s []rune, coin int) func(name string) int {
	return func(name string) int {
		var reg strings.Builder
		length := len(s)
		if length == 0 {
			return 0
		}

		for i, str := range s {
			reg.WriteString(string(str))
			if i+1 < length {
				reg.WriteString("|")
			}
		}

		strLen := len(regexp.MustCompile(reg.String()).FindAllStringIndex(name, -1))

		return strLen * coin
	}
}
