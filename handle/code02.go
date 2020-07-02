package handle

import (
	"encoding/json"
	"fmt"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 别人的
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		curr.Next = new(ListNode)
		curr = curr.Next
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		curr.Val = carry % 10
		carry /= 10
	}
	return dummy.Next
}

// ok mine 自己的
/*
	两个链表相加
	2,4,3
	5,6,4,9

	2+5=7,  4+6(需要进一位)=0, 3+4+1(进一位)=8, 0+9=9 ...
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	l3.Next = new(ListNode)
	prev := l3.Next
	extra := 0
	for {
		val := 0
		if l1 != nil {
			val += l1.Val
		}

		if l2 != nil {
			val += l2.Val
		}

		val += extra

		if val >= 10 {
			val = val - 10
			extra = 1
		} else {
			extra = 0
		}

		prev.Val = val
		if (l1 == nil || l1.Next == nil) && (l2 == nil || l2.Next == nil) {
			if extra == 1 {
				prev.Next = new(ListNode)
				prev = prev.Next
				prev.Val = extra
			}
			break
		}

		prev.Next = new(ListNode)
		prev = prev.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return l3.Next
}

func ParseListNode(list []int) *ListNode {
	node := &ListNode{list[0], nil}
	if len(list) == 1 {
		return node
	}
	node.Next = &ListNode{}
	temp := node.Next

	for i := 1; i < len(list); i++ {
		temp.Val = list[i]
		if i == len(list)-1 {
			break
		}
		temp.Next = new(ListNode)
		temp = temp.Next
	}
	return node
}

func AddTwoNumbers5(l1 *ListNode, l2 *ListNode) *ListNode {
	//
	var expand = func(l *ListNode) []int {
		var arr []int
		for {
			arr = append(arr, l.Val)

			if l.Next == nil {
				return arr
			}
			l = l.Next
		}
	}

	list1 := expand(l1)
	list2 := expand(l2)

	length := len(list1)
	if len(list2) > length {
		length = len(list2)
	}

	var clist []int
	j := 0
	for i := 0; i < length; i++ {
		a := 0
		b := 0
		if i < len(list1) {
			a = list1[i]
		}
		if i < len(list2) {
			b = list2[i]
		}
		t := a + b + j
		if t < 10 {
			clist = append(clist, t)
			j = 0
		} else {
			clist = append(clist, t-10)
			j = 1
		}
	}
	if j == 1 {
		clist = append(clist, j)
	}

	fmt.Println("clist", clist)

	l3 := &ListNode{clist[0], nil}
	if len(clist) == 1 {
		return l3
	}

	l3.Next = &ListNode{}
	temp := l3.Next
	for i := 1; i < len(clist); i++ {
		temp.Val = clist[i]
		fmt.Println("----", temp.Val)
		if i < len(clist)-1 {
			temp.Next = new(ListNode)
			temp = temp.Next
		}
	}

	return l3
}

func AddTwoNumbers4(l1 *ListNode, l2 *ListNode) *ListNode {
	//
	var expand = func(l *ListNode) []int {
		var arr []int
		for {
			arr = append(arr, l.Val)

			if l.Next == nil {
				return arr
			}
			l = l.Next
		}
	}

	list1 := expand(l1)
	list2 := expand(l2)

	var str2Int = func(s string) int {
		n, _ := strconv.Atoi(s)
		return n
	}

	var arr2Int = func(a []int) int {
		fmt.Println("list:", a)
		var num string
		for i := len(a) - 1; i >= 0; i-- {
			fmt.Println("--", a[i])
			num = fmt.Sprintf("%s%d", num, a[i])
		}
		fmt.Println("num", num)
		return str2Int(num)
	}

	var left []int
	var a, b int
	if len(list1) <= len(list2) {
		a = arr2Int(list1)
		b = arr2Int(list2[:len(list1)])
		left = list2[len(list1):]
	} else {
		a = arr2Int(list1[:len(list2)])
		b = arr2Int(list2)
		left = list1[len(list2):]
	}

	c := fmt.Sprintf("%d", a+b)
	fmt.Println("===c", c, "left", left)
	l3 := &ListNode{str2Int(string(c[len(c)-1])), nil}
	if len(c) == 1 && len(left) == 0 {
		return l3
	}
	fmt.Println("l3", l3.Val)

	l3.Next = &ListNode{}
	temp := l3.Next
	for i := len(c) - 2; i >= 0; i-- {
		if temp == nil {
			temp = new(ListNode)
		}
		temp.Val = str2Int(string(c[i]))
		fmt.Println("----", temp.Val)
		if i > 0 {
			temp.Next = new(ListNode)
			temp = temp.Next
		}
	}

	bbb, _ := json.Marshal(l3)
	fmt.Println("=======", string(bbb), temp)
	if len(left) == 0 {
		return l3
	}

	if temp == nil {
		l3.Next = new(ListNode)
		temp = l3.Next
	}

	for i := 0; i < len(left); i++ {
		temp.Val = left[i]
		if i < len(left)-1 {
			temp.Next = new(ListNode)
			temp = temp.Next
		}

	}

	bbb, _ = json.Marshal(l3)
	fmt.Println("=======", string(bbb))
	return l3
}

/*
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
// [1,0,9]
// [5,7,8]
func AddTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	str := ""
	strorin := ""
	for {
		str = fmt.Sprintf("%d%s", l1.Val, str)
		strorin = fmt.Sprintf("%s%d", strorin, l1.Val)

		if l1.Next == nil {
			break
		}
		l1 = l1.Next
	}
	//a, _ := strconv.Atoi(str)

	str2 := ""
	strorin2 := ""
	for {
		str2 = fmt.Sprintf("%d%s", l2.Val, str2)
		strorin2 = fmt.Sprintf("%s%d", strorin2, l2.Val)

		if l2.Next == nil {
			break
		}
		l2 = l2.Next
	}

	fmt.Println("origin:", strorin, strorin2)
	fmt.Println("len(str)", len(str), len(str2))
	var c string

	left := ""
	a := 0
	b := 0
	if len(strorin) <= len(strorin2) {
		a, _ = strconv.Atoi(str)
		b, _ = strconv.Atoi(str2[:len(str)])
		left = strorin2[len(strorin):]

	} else {
		a, _ = strconv.Atoi(str[:len(str2)])
		b, _ = strconv.Atoi(str2)
		left = strorin[len(strorin2):]

	}

	c = string(fmt.Sprintf("%d", a+b))

	var l3 ListNode
	v, _ := strconv.Atoi(string(c[len(c)-1]))
	l3.Val = v

	if len(c) == 1 {
		return &l3
	}

	fmt.Println(a, b, c)
	l3.Next = &ListNode{}
	temp := l3.Next

	for i := len(c) - 2; i >= 0; i-- {
		v, _ := strconv.Atoi(string(c[i]))
		temp.Val = v
		if i == 0 {
			break
		}
		temp.Next = &ListNode{
			Val: v,
		}
		temp = temp.Next
	}

	bbb, _ := json.Marshal(l3)
	fmt.Println(string(bbb))
	for i := 0; i < len(left); i++ {
		v, _ := strconv.Atoi(string(left[i]))
		temp.Next = &ListNode{
			Val: v,
		}
		temp = temp.Next
	}
	return &l3
}

func AddTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	a := 0
	index := 1
	for {
		if index > 1 && l1.Val == 0 {
			l1.Val = 1
		}
		fmt.Println("l1.va", l1.Val)
		a += index * l1.Val
		if l1.Next == nil {
			break
		}
		l1 = l1.Next
		index *= 10
	}
	fmt.Println("a", a)
	index = 1
	b := 0
	for {
		if index > 1 && l2.Val == 0 {
			l2.Val = 1
		}

		b += index * l2.Val

		if l2.Next == nil {
			break
		}
		l2 = l2.Next
		index *= 10
	}
	fmt.Println("b", b)
	c := string(fmt.Sprintf("%d", a+b))
	fmt.Println("c", c)
	var l3 ListNode
	v, _ := strconv.Atoi(string(c[len(c)-1]))
	l3.Val = v

	if len(c) == 1 {
		return &l3
	}

	l3.Next = &ListNode{}
	temp := l3.Next
	for i := len(c) - 2; i >= 0; i-- {
		v, _ := strconv.Atoi(string(c[i]))
		temp.Val = v
		if i == 0 {
			break
		}
		temp.Next = &ListNode{
			Val: v,
		}
		temp = temp.Next
	}
	return &l3
}
