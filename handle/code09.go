package handle

import "fmt"

// 回文数
// 反转一半然后进行比较, 何时到一半只要判断x>y 就好了,然后相等的话 x==y/10
func isPalindrome(x int) bool {
	// 1234321
	// 10
	//i := 0
	// 需要排除 结尾是0的,因为如果结尾是0 反转之后肯定是错的,例如 100 -> 001 0不能开头
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	y := 0
	for x > y {
		y = y*10 + x%10
		fmt.Println("x,y", x, y)
		x /= 10
	}
	return x == y || x == y/10
}

// 自己的
// 直接反转数字 比较结果是否相等
//从leetcode上看, 这个直接反转然后比较是否一致在时间和内存上都要好一点,不知道为啥, 但是按到理应该上面的算法要好一点
func isPalindrome2(x int) bool {

	if x < 0 {
		return false
	}
	o := x
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x /= 10
	}
	return o == y
}

func IsPalindrome(x int) bool {
	return isPalindrome(x)
}
