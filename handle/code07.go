package handle

import (
	"fmt"
	"strconv"
)

// 7. 整数反转

/*

给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 自己的  字符串转一次
func Reverse1(x int) int {
	//a :=
	var res string
	var p bool
	if x < 0 {
		p = true
		x = -x
	}
	for {
		left := x % 10
		res += fmt.Sprintf("%d", left)
		//fmt.Println("res:", res)
		if x < 10 {
			break
		}
		x = x / 10
	}

	//fmt.Println("string:", res)
	v, _ := strconv.Atoi(res)
	b := 1024 * 1024 * 1024 * 2
	if v > b-1 || v < -b {
		return 0
	}
	//fmt.Println("intatoi", v)
	if p {
		v = -v
	}

	//fmt.Println("intatoi", v)
	return v
}

// 别人的
func Reverse(x int) int {
	y := 0
	b := 1 << 31
	for x != 0 {
		y = y*10 + x%10
		if y > -b || y > b-1 {
			return 0
		}
		x /= 10
	}
	return y
}

/* 作者：elliotxx
链接：https://leetcode-cn.com/problems/reverse-integer/solution/0ms-11-xing-dai-ma-go-shi-xian-by-elliotxx/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。 */
