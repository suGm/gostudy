package main

import "fmt"

// 必须制定存放元素的类型和容量（长度）
// 数组的长度是数组类型的一部分
func main() {

	//var a1 [3]bool
	//var a2 [3]bool
	//
	//a1 = [3]bool{true,true,false}
	//a2 = [3]bool{false,true,true}
	//
	//fmt.Printf("%T\t %T\n", a1, a2)
	//fmt.Println(a1 == a2)
	//
	//// 初始化数组
	//a10 := [...]int{1,2,3,4,5,6,7}
	//// 根据索引初始化
	//a5  := [5]int{0:1,4:2}
	//fmt.Println(a10, a5)
	//
	//// 数组遍历
	//
	//// 根据索引遍历
	//for i := 0; i < len(a10); i++ {
	//	fmt.Println(a10[i])
	//}
	//
	//// 多维数组
	//// [[1,2][3,4][5,6]]
	//var a11 [3][2]int
	//a11 = [3][2]int{
	//	[2]int{1,2},
	//	[2]int{3,4},
	//	[2]int{5,6},
	//}
	//
	//// 多维数组的遍历
	//for _,v := range a11{
	//	fmt.Println(v)
	//	for _,v2 := range v{
	//		fmt.Println(v2)
	//	}
	//}
	//
	//// 数组是值类型 不是引用类型
	//b1 := [3]int{1,2,3}
	//b2 := b1
	//b2[0] = 100//{100,2,3}
	//fmt.Println(b1,b2)
	//
	//a12 := [5]int{1,2,5,7,8}
	//count := 0
	//for _,v3 := range a12 {
	//	count += v3
	//}
	//fmt.Println(count)
	//
	//a13 := [2][2]int{}
	//x := 0
	//for i,v4 := range a12 {
	//	for j := i+1; j < len(a12); j++ {
	//		if v4+a12[j] == 9 {
	//			a13[x][0] = v4
	//			a13[x][1] = a12[j]
	//			x++
	//		}
	//	}
	//}
	//fmt.Println(a13)

	var list = ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			},
		},
	}
	FindKthToTail(&list, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here

	slow, fast := pHead, pHead

	for k > 0 && fast != nil {
		fast = fast.Next
		k--
	}

	if k > 0 {
		return nil
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	for slow != nil {
		fmt.Println(slow.Val)
		slow = slow.Next
	}

	return slow
}
