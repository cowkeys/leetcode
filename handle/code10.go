package handle

import "fmt"

// 正则表达式 '*' '.'
/*

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false

*/
func IsMatch(s string, p string) bool {
	if p == "" || s == "" {
		return false
	}

	if s == p {
		return true
	}

	// abbcdddd
	// a.*cd*
	var prev byte
	i, j := 0, 0
	for i = 0; i < len(s); i++ {
		if j > len(p)-1 {
			break
		}
		// 如果相等
		//fmt.Println("s[i] , p[j] ", string(s[i]), string(p[j]))
		if s[i] == p[j] {
			//fmt.Printf("s[i]==p[j] i:%d,j:%d,s[%d]:%s,p[%d]:%s, prev:%s \n", i, j, i, string(s[i]), j, string(p[j]), string(prev))
			j++
			prev = s[i]
			continue
		}
		// 46 : .
		if p[j] == 46 {
			j++
			prev = s[i]
			continue
		}
		// 42 : *
		if p[j] == 42 {
			//fmt.Printf("p[j]==42 i:%d,j:%d,s[i]:%s,p[j]:%s, prev:%s \n", i, j, string(s[i]), string(p[j]), string(prev))
			if s[i] != prev {
				j++
				prev = s[i]
				continue
			}

			for s[i] == prev {
				i++
				if i > len(s)-1 {
					i--
					break
				}
				//fmt.Printf("for i++ s[i]:%s,prev:%s i:%d j:%d\n", string(s[i]), string(prev), i, j)
			}

			j++
			//fmt.Println("in i,j", i, j)
			continue
		}

		if j+1 < len(p) {
			//fmt.Println("xx", p[j+1])
			if p[j+1] == 42 {
				j += 2
				//prev = s[i]
				continue
			}
		}
		return false
	}
	//fmt.Println(s, p, i, j)

	if j < len(p)-1 {
		return false
	}
	fmt.Println("---", s, p)
	//fmt.Println(s, "xxxx true")
	return true
}
